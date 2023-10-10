package main

import (
	"log"
	"senpainikolay/pad-sem7/accident-reporting-service/config"
	"senpainikolay/pad-sem7/accident-reporting-service/internal/controller"
	"senpainikolay/pad-sem7/accident-reporting-service/internal/models"
	"senpainikolay/pad-sem7/accident-reporting-service/internal/repository"
	"senpainikolay/pad-sem7/accident-reporting-service/internal/service"
	postgres "senpainikolay/pad-sem7/accident-reporting-service/pkg"

	"gorm.io/gorm"
)

var db *gorm.DB
var cnf config.Config
var yamlFilePath = "config/config.yaml"

func main() {

	accRepo := repository.NewAccidentRepository(db)
	accService := service.NewAccidentService(accRepo)

	log.Printf("starting gRPC API server...\n")
	controller.Serve(accService, cnf.ServicePort)

}

func init() {
	cnf = config.NewConfig(yamlFilePath)
	db = postgres.NewDBConnection(cnf)
	err := db.AutoMigrate(models.AccidentModel{})
	if err != nil {
		log.Fatalf("failed to migrate user model\n")
	}

	sqlCommands := []string{
		"CREATE EXTENSION IF NOT EXISTS postgis;",
		"CREATE EXTENSION IF NOT EXISTS postgis_topology;",
	}

	for _, sqlCommand := range sqlCommands {
		if err := db.Exec(sqlCommand).Error; err != nil {
			log.Fatalf(err.Error())
		}
	}

}
