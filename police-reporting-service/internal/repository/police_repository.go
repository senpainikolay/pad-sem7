package repository

import (
	"context"
	"errors"
	"time"

	"senpainikolay/pad-sem7/police-reporting-service/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const CONFIRMATION_THRESHOLD = 2

type PoliceRepository struct {
	dbClient *mongo.Client
	dbName   string
}

func NewPoliceRepository(dbClient *mongo.Client, dbName string) *PoliceRepository {
	return &PoliceRepository{
		dbClient: dbClient,
		dbName:   dbName,
	}
}

func (repo *PoliceRepository) FetchPolice(usrInfo models.UserGeoInfo) (models.PoliceGeoInfoResponse, error) {

	cityCol := repo.dbClient.Database(repo.dbName).Collection(usrInfo.City)

	queryFilterUsrCoordsInfo := bson.D{{"type", "Point"}, {"coordinates", []float64{usrInfo.UserLong, usrInfo.UserLat}}}
	filter := bson.D{
		{"location", bson.D{
			{"$near", bson.D{
				{"$geometry", queryFilterUsrCoordsInfo},
				{"$maxDistance", usrInfo.ZoomIndex * 10000}, // 10k meters * zoom_index
			}},
		}},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filterCursor, err := cityCol.Find(ctx, filter)
	if err != nil {
		return models.PoliceGeoInfoResponse{}, err
	}

	defer filterCursor.Close(ctx)

	var policeInfoFetch models.PoliceGeoInfoResponse

	for filterCursor.Next(ctx) {

		var polEntity models.PoliceEntity

		if err = filterCursor.Decode(&polEntity); err != nil {
			return models.PoliceGeoInfoResponse{}, err
		}

		if err = updateConfirmationBool(&polEntity, cityCol); err != nil {
			return models.PoliceGeoInfoResponse{}, err
		}

		policeInfoFetch.Data = append(policeInfoFetch.Data, models.PoliceInfo{
			PolLong:                  polEntity.Location.Coordinates[0],
			PolLat:                   polEntity.Location.Coordinates[1],
			ConfirmationNotification: polEntity.ConfirmationNotification,
			ConfirmedBy:              polEntity.ConfirmedBy,
		})

	}

	return policeInfoFetch, nil

}

func updateConfirmationBool(pol *models.PoliceEntity, col *mongo.Collection) error {
	durationInSeconds := time.Now().Unix() - pol.UpdatedAt.Unix()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if durationInSeconds > 300 && pol.ConfirmationNotification != true { // > 5 minutes => confirming the locations again
		_, err := col.UpdateOne(
			ctx,
			bson.M{"_id": pol.ID},
			bson.D{
				{"$set", bson.D{{"confirmation_notification", true},
					{"updated_at", time.Now()},
					{"confirmed_by", 0},
				}},
			},
		)
		if err != nil {
			return err
		}
		pol.ConfirmationNotification = true
		pol.UpdatedAt = time.Now()
	}
	return nil
}

func (repo *PoliceRepository) PostPolice(polM models.PolicePostInfo) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	createIndexBool, err := repo.checkIfToCreateIndex(&ctx, polM.City)
	if err != nil {
		return err
	}
	cityCol := repo.dbClient.Database(repo.dbName).Collection(polM.City)

	if createIndexBool {
		indexOpts := options.CreateIndexes().SetMaxTime(time.Second * 10)
		pointIndexModel := mongo.IndexModel{
			Keys: bson.M{"location": "2dsphere"},
		}

		pointIndexes := cityCol.Indexes()

		_, err := pointIndexes.CreateOne(ctx, pointIndexModel, indexOpts)
		if err != nil {
			return err
		}
	}

	polEntity := models.PoliceEntity{
		Location:                 models.NewPoint(polM.PolLong, polM.PolLat),
		ConfirmationNotification: true,
		ConfirmedBy:              0,
		CreatedAt:                time.Now(),
		UpdatedAt:                time.Now(),
	}

	_, insertErr := cityCol.InsertOne(ctx, polEntity)
	if insertErr != nil {
		return insertErr
	}
	return nil
}

func (repo *PoliceRepository) checkIfToCreateIndex(c *context.Context, cityName string) (bool, error) {
	cityCols, err := repo.dbClient.Database(repo.dbName).ListCollectionNames(*c, bson.M{})
	if err != nil {
		return false, err
	}

	for _, city := range cityCols {
		if city == cityName {
			return false, nil
		}
	}
	return true, nil

}

func (repo *PoliceRepository) UpdatePoliceCoordsConfirmation(polInfo models.PolicePostInfo) error {
	cityCol := repo.dbClient.Database(repo.dbName).Collection(polInfo.City)

	filter := bson.D{
		{"location.coordinates", []float64{polInfo.PolLong, polInfo.PolLat}},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	polRes := cityCol.FindOne(ctx, filter)

	var polEntity models.PoliceEntity
	err := polRes.Decode(&polEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("coordinates not found!!")
		}
		return err
	}

	if polEntity.ConfirmedBy > CONFIRMATION_THRESHOLD && polEntity.ConfirmationNotification == true {
		_, err = cityCol.UpdateOne(
			ctx,
			bson.M{"_id": polEntity.ID},
			bson.D{{"$set", bson.D{
				{"confirmation_notification", false},
				{"updated_at", time.Now()},
			}}},
		)
		if err != nil {
			return err
		}
	}

	if polEntity.ConfirmedBy <= CONFIRMATION_THRESHOLD {

		_, err = cityCol.UpdateOne(
			ctx,
			bson.M{"_id": polEntity.ID},
			bson.D{
				{"$set", bson.D{
					{"updated_at", time.Now()},
					{"confirmed_by", polEntity.ConfirmedBy + 1},
				}}},
		)
		if err != nil {
			return err
		}
	}
	return nil

}

func (repo *PoliceRepository) DeletePoliceCoords(polInfo models.PolicePostInfo) error {

	cityCol := repo.dbClient.Database(repo.dbName).Collection(polInfo.City)

	filter := bson.D{
		{"location.coordinates", []float64{polInfo.PolLong, polInfo.PolLat}},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	polRes := cityCol.FindOne(ctx, filter)

	var polEntity models.PoliceEntity
	err := polRes.Decode(&polEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("coordinates not found!!")
		}
		return err
	}

	_, err = cityCol.DeleteOne(
		ctx,
		bson.M{"_id": polEntity.ID},
	)

	return err

}
