package main

import (
	"log"
	"senpainikolay/pad-sem7/police-reporting-service/internal/controller"
	"senpainikolay/pad-sem7/police-reporting-service/internal/repository"
	"senpainikolay/pad-sem7/police-reporting-service/internal/service"
	mongodbp "senpainikolay/pad-sem7/police-reporting-service/pkg"
)

func main() {
	dbName := "test_db"
	MONGO_CLIENT := mongodbp.NewDBConnection()

	policeRepo := repository.NewPoliceRepository(MONGO_CLIENT, dbName)
	policeService := service.NewPoliceService(policeRepo)

	log.Printf("starting gRPC API server...\n")
	controller.Serve(policeService, ":6666")

}
