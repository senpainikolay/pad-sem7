package service

import (
	"fmt"
	"log"
	"senpainikolay/pad-sem7/accident-reporting-service/internal/models"
	policeserviceclient "senpainikolay/pad-sem7/accident-reporting-service/internal/police-service-client"
	"time"
)

type IAccidentRepository interface {
	FetchAccidents(models.UserGeoInfo) (models.AccidentGeoInfoResponse, error)
	PostAccident(*models.AccidentModel) error
	GetByPos(float64, float64) (models.AccidentModel, error)
	DeleteByPos(float64, float64) error
	UpdateAccConfirmationNot(uint, bool) error
	UpdateAccConfirmationIndex(uint, int) error
}

const THRESHOLD_TIMEOUT_SERVICE_CALL = 900

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

func (svc *AccidentService) ConfirmAccident(confInfo models.ConfirmationAccidentInfo) error {

	if confInfo.PoliceConfirmation == true || confInfo.AccidentConfirmation == false {
		return svc.accidentRepo.DeleteByPos(confInfo.Long, confInfo.Lat)
	}

	acc, err := svc.accidentRepo.GetByPos(confInfo.Long, confInfo.Lat)
	if err != nil {
		return err
	}

	if (time.Now().Unix()-acc.CreatedAt.Unix()) > THRESHOLD_TIMEOUT_SERVICE_CALL && !confInfo.PoliceConfirmation && confInfo.AccidentConfirmation {
		var data models.ExernalServiceData
		data.CarsInvolved = int(acc.CarInvolved)
		data.City = acc.City
		data.StreetName = acc.StreetName
		err = policeserviceclient.InformExternalService(data, confInfo.Long, confInfo.Lat)
		if err != nil {
			log.Println("ERROR ON CALLING THE EXTERNAL SERVICE: %v", err.Error())
		}
		return svc.accidentRepo.DeleteByPos(confInfo.Long, confInfo.Lat)
	}

	if acc.ConfirmedBy > 2 {
		return svc.accidentRepo.UpdateAccConfirmationNot(acc.ID, false)
	}

	return svc.accidentRepo.UpdateAccConfirmationIndex(acc.ID, acc.ConfirmedBy)

}
