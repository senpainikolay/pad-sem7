package models

import "gorm.io/gorm"

type AccidentModel struct {
	*gorm.Model
	ConfirmationAccidentNotification bool   `json:"confirmation_accident_notification"`
	ConfirmationPoliceNotification   bool   `json:"confirmation_police_notification"`
	ConfirmedBy                      int    `json:"confirmed_by"`
	Coordinates                      string `gorm:"type:geometry(Point)`
	City                             string `json:"city"`
	StreetName                       string `json:"street_name"`
	CarInvolved                      int64  `json:"cars_involved"`
}

type UserGeoInfo struct {
	UserLong  float64 `json:"user_long"`
	UserLat   float64 `json:"user_lat"`
	ZoomIndex int64   `json:"zoom_index"`
	City      string  `json:"city"`
}

type AccidentInfo struct {
	Long                             float64 `json:"long"`
	Lat                              float64 `json:"lat"`
	ConfirmedBy                      int64   `json:"confirmed_by"`
	ConfirmationAccidentNotification bool    `json:"confirmation_accident_notification"`
	ConfirmationPoliceNotification   bool    `json:"confirmation_police_notification"`
}

type AccidentPostInfo struct {
	Long        float64 `json:"long"`
	Lat         float64 `json:"lat"`
	CarInvolved int64   `json:"cars_involved"`
	City        string  `json:"city"`
	StreetName  string  `json:"street_name"`
}

type ConfirmationAccidentInfo struct {
	Long                 float64 `json:"long"`
	Lat                  float64 `json:"lat"`
	PoliceConfirmation   bool    `json:"police_confimation"`
	AccidentConfirmation bool    `json:"accident_confimation"`
}

type AccidentGeoInfoResponse struct {
	Data []AccidentInfo `json:"data"`
}
