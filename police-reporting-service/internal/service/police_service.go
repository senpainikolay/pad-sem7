package service

import "senpainikolay/pad-sem7/police-reporting-service/internal/models"

type IPoliceRepository interface {
	PostPolice(models.PolicePostInfo) error
	FetchPolice(models.UserGeoInfo) (models.PoliceGeoInfoResponse, error)
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

func (svc *PoliceService) ConfirmPolice(models.ConfirmationPoliceInfo) error {
	return nil
}
