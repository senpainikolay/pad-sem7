package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Location struct {
	Type        string    `bson:"type"`
	Coordinates []float64 `bson:"coordinates"`
}

func NewPoint(long, lat float64) Location {
	return Location{
		"Point",
		[]float64{long, lat},
	}
}

type PoliceEntity struct {
	ID                       primitive.ObjectID `bson:"_id,omitempty"`
	Location                 Location           `bson:"location"`
	ConfirmationNotification bool               `bson:"confirmation_notification"`
	ConfirmedBy              int                `bson:"confirmed_by"`
	CreatedAt                time.Time          `bson:"created_at"`
	UpdatedAt                time.Time          `bson:"updated_at"`
}

type UserGeoInfo struct {
	UserLong  float64 `json:"user_long"`
	UserLat   float64 `json:"user_lat"`
	ZoomIndex int64   `json:"zoom_index"`
	City      string  `json:"city"`
}

type PolicePostInfo struct {
	PolLong float64 `json:"pol_long"`
	PolLat  float64 `json:"pol_lat"`
	City    string  `json:"city"`
}

type ConfirmationPoliceInfo struct {
	PoliceInfo   PolicePostInfo
	Confirmation bool `json:"confimation"`
}

type PoliceInfo struct {
	PolLong                  float64 `json:"pol_long"`
	PolLat                   float64 `json:"pol_lat"`
	ConfirmationNotification bool    `json:"confirmation_notification"`
	ConfirmedBy              int     `json:"confirmed_by"`
}

type PoliceGeoInfoResponse struct {
	Data []PoliceInfo `json:"data"`
}
