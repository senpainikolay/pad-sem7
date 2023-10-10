package service

import (
	"fmt"
	"senpainikolay/pad-sem7/accident-reporting-service/internal/models"
)

type IAccidentRepository interface {
	FetchAccidents(models.UserGeoInfo) (models.AccidentGeoInfoResponse, error)
	PostAccident(*models.AccidentModel) error
}

type AccidentService struct {
	accidentRepo IAccidentRepository
}

func NewAccidentService(accRepo IAccidentRepository) *AccidentService {
	return &AccidentService{
		accidentRepo: accRepo,
	}
}

func (svc *AccidentService) Fetch(usrInfo models.UserGeoInfo) (models.AccidentGeoInfoResponse, error) {
	return svc.accidentRepo.FetchAccidents(usrInfo)
}
func (svc *AccidentService) PostAccident(accInfo models.AccidentPostInfo) error {
	return svc.accidentRepo.PostAccident(&models.AccidentModel{
		ConfirmationAccidentNotification: true,
		ConfirmationPoliceNotification:   false,
		ConfirmedBy:                      0,
		City:                             accInfo.City,
		StreetName:                       accInfo.StreetName,
		CarInvolved:                      accInfo.CarInvolved,
		Coordinates:                      fmt.Sprintf("POINT(%f %f)", accInfo.Long, accInfo.Lat),
	})
}
