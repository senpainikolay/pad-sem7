package models

type HealthInfo struct {
	Ready    bool   `json:"ready"`
	Database string `json"database"`
	Load     bool   `json:"load"`
}
