package service

import (
	"senpainikolay/pad-sem7/police-reporting-service/internal/models"

	"golang.org/x/time/rate"
)

type IPoliceRepository interface {
	PostPolice(models.PolicePostInfo) error
	FetchPolice(models.UserGeoInfo) (models.PoliceGeoInfoResponse, error)
	UpdatePoliceCoordsConfirmation(models.PolicePostInfo) error
	DeletePoliceCoords(models.PolicePostInfo) error
	Ping() error
}

type PoliceService struct {
	policeRepo IPoliceRepository
}

func NewPoliceService(policeRepo IPoliceRepository) *PoliceService {
	return &PoliceService{
		policeRepo: policeRepo,
	}
}

func (svc *PoliceService) PostPolice(policeInfo models.PolicePostInfo) error {
	return svc.policeRepo.PostPolice(policeInfo)
}

func (svc *PoliceService) Fetch(usrInfo models.UserGeoInfo) (models.PoliceGeoInfoResponse, error) {
	return svc.policeRepo.FetchPolice(usrInfo)
}

func (svc *PoliceService) ConfirmPolice(confirmInfo models.ConfirmationPoliceInfo) error {

	if confirmInfo.Confirmation == true {
		return svc.policeRepo.UpdatePoliceCoordsConfirmation(confirmInfo.PoliceInfo)
	}

	return svc.policeRepo.DeletePoliceCoords(confirmInfo.PoliceInfo)
}

func (svc *PoliceService) CheckHealth(rl *rate.Limiter) models.HealthInfo {
	healthInfo := models.HealthInfo{
		Ready:    true,
		Database: "connected",
		Load:     false,
	}

	err := svc.policeRepo.Ping()
	if err != nil {
		healthInfo.Database = "disconnected"
		healthInfo.Ready = false
	}

	if !rl.Allow() {
		healthInfo.Load = true
		healthInfo.Ready = false

	}
	return healthInfo

}
