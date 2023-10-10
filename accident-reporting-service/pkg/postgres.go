package postgres

import (
	"fmt"
	"log"
	"senpainikolay/pad-sem7/accident-reporting-service/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBConnection(cnf config.Config) *gorm.DB {
	dsn := fmt.Sprintf(`host=%s
						user=%s 
						password=%s 
						dbname=%s 
						port=%s 
						sslmode=disable`,
		cnf.PostgresHost,
		cnf.UserDB,
		cnf.PostgresPW,
		cnf.PostgresDBName,
		cnf.PostgresPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed gorm connection: %v\n", err)
	}

	log.Printf("successfully connect to postgres db...\n")
	return db
}
