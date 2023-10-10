package models

type PoliceInfo struct {
	Long        float64
	Lat         float64
	ConfirmedBy int
}

type ExernalServiceData struct {
	CarsInvolved int
	City         string
	StreetName   string
	NearbyPolice []PoliceInfo
}
