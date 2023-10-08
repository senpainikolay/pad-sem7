package models

import "gorm.io/gorm"

type AccidentModel struct {
	*gorm.Model
	ConfirmationAccidentNotification bool   `json:"confirmation_accident_notification"`
	ConfirmationPoliceNotification   bool   `json:"confirmation_police_notification"`
	ConfirmedBy                      int    `json:"confirmed_by"`
	Coordinates                      string `gorm:"type:geometry(Point)"`
	City                             string `json:"city"`
	StreetName                       string `json:"street_name"`
}

type UserGeoInfo struct {
	UserLong  float64 `json:"user_long"`
	UserLat   float64 `json:"user_lat"`
	ZoomIndex int64   `json:"zoom_index"`
	City      string  `json:"city"`
}

type AccidentInfo struct {
	PolLong float64 `json:"pol_long"`
	PolLat  float64 `json:"pol_lat"`
	City    string  `json:"city"`
}

type AccidentPostInfo struct {
	AccidentInfo
	StreetName string `json:"street_name"`
}

type ConfirmationPoliceInfo struct {
	AccidentInfo         AccidentInfo
	PoliceConfirmation   bool `json:"police_confimation"`
	AccidentConfirmation bool `json:"accident_confimation"`
}

type PoliceGeoInfoResponse struct {
	Data []AccidentPostInfo `json:"data"`
}
