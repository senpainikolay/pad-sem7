package repository

import (
	"fmt"
	"log"
	"senpainikolay/pad-sem7/accident-reporting-service/internal/models"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type AccidentRepository struct {
	dbClient *gorm.DB
}

func NewAccidentRepository(dbClient *gorm.DB) *AccidentRepository {
	return &AccidentRepository{
		dbClient: dbClient,
	}
}

func (repo *AccidentRepository) PostAccident(acc *models.AccidentModel) error {
	err := repo.dbClient.Debug().
		Model(models.AccidentModel{}).
		Create(acc).Error
	if err != nil {
		log.Printf("failed to insest user in database: %v\n", err)
		return err
	}
	return nil
}

func (repo *AccidentRepository) FetchAccidents(usrInfo models.UserGeoInfo) (models.AccidentGeoInfoResponse, error) {

	var accidents []models.AccidentModel

	query := `
        SELECT *
        FROM accident_models
        WHERE ST_DWithin(
            ST_GeomFromText('POINT(` + fmt.Sprintf("%.4f %.4f ", usrInfo.UserLong, usrInfo.UserLat) + `)', 4326),
            "coordinates"::geometry,
            ` + fmt.Sprintf("%d", usrInfo.ZoomIndex*10000) + `, 
			true
        )
    `

	if err := repo.dbClient.Raw(query).Scan(&accidents).Error; err != nil {
		return models.AccidentGeoInfoResponse{}, err
	}

	var resp models.AccidentGeoInfoResponse

	for _, accModel := range accidents {
		long, lat := decodePointString(accModel.Coordinates)

		resp.Data = append(resp.Data, models.AccidentInfo{
			Long:                             long,
			Lat:                              lat,
			ConfirmedBy:                      int64(accModel.ConfirmedBy),
			ConfirmationAccidentNotification: accModel.ConfirmationAccidentNotification,
			ConfirmationPoliceNotification:   accModel.ConfirmationPoliceNotification,
		})
	}
	return resp, nil

}

func decodePointString(pointString string) (float64, float64) {
	pointString = strings.TrimPrefix(pointString, "POINT(")
	pointString = strings.TrimSuffix(pointString, ")")

	coordinates := strings.Split(pointString, " ")
	if len(coordinates) != 2 {
		return 0.0, 0.0
	}
	longitude, _ := strconv.ParseFloat(coordinates[0], 64)
	latitude, _ := strconv.ParseFloat(coordinates[1], 64)

	return longitude, latitude
}
