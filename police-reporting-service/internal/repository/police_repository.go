package repository

import (
	"context"
	"time"

	"senpainikolay/pad-sem7/police-reporting-service/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

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
		{"$near", bson.D{
			{"$geometry", queryFilterUsrCoordsInfo},
			{"$maxDistance", 2000}, // meters
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
			PolLong:                  polEntity.Coordonates[0],
			PolLat:                   polEntity.Coordonates[1],
			ConfirmationNotification: polEntity.ConfirmationNotification,
			ConfirmedBy:              polEntity.ConfirmedBy,
		})

	}

	return policeInfoFetch, nil

}

func updateConfirmationBool(pol *models.PoliceEntity, col *mongo.Collection) error {
	durationInSeconds := pol.UpdatedAt.Unix() - pol.CreatedAt.Unix()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if durationInSeconds > 300 && pol.ConfirmationNotification != true { // > 5 minutes => confirming the locations again
		_, err := col.UpdateOne(
			ctx,
			bson.M{"_id": pol.ID},
			bson.D{
				{"$set", bson.D{{"confirmation_notification", true},
					{"updated_at", time.Now()},
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

	cityCol := repo.dbClient.Database(repo.dbName).Collection(polM.City)

	polEntity := models.PoliceEntity{
		Coordonates:              []float64{polM.PolLong, polM.PolLat},
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
