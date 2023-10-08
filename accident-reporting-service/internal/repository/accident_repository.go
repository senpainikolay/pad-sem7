package repository

import "gorm.io/gorm"

type AccidentRepository struct {
	dbClient *gorm.DB
}

func NewAccidentRepository(dbClient *gorm.DB) *AccidentRepository {
	return &AccidentRepository{
		dbClient: dbClient,
	}
}
