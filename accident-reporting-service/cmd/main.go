package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"senpainikolay/pad-sem7/accident-reporting-service/config"
	"senpainikolay/pad-sem7/accident-reporting-service/internal/controller"
	"senpainikolay/pad-sem7/accident-reporting-service/internal/models"
	"senpainikolay/pad-sem7/accident-reporting-service/internal/repository"
	"senpainikolay/pad-sem7/accident-reporting-service/internal/service"
	postgres "senpainikolay/pad-sem7/accident-reporting-service/pkg"
	"syscall"

	"gorm.io/gorm"
)

var db *gorm.DB
var cnf config.Config
var yamlFilePath = "config/config.yaml"

func main() {

	accRepo := repository.NewAccidentRepository(db)
	accService := service.NewAccidentService(accRepo)

	serviceDiscoveryReq("register", "localhost:6667", "POST")

	// in case of force stop
	go signalUnregisterThread()

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

func serviceDiscoveryReq(route, serviceUrl, method string) {

	url := fmt.Sprintf("http://localhost:8000/services/%s", route)

	requestData := map[string]string{
		"service_type": "accident-reporting",
		"service_url":  serviceUrl,
	}

	requestDataBytes, err := json.Marshal(requestData)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestDataBytes))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return
	} else {
		log.Printf("Request not succesful, code: %v", resp.StatusCode)
	}
}

func signalUnregisterThread() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	log.Printf("FORCE EXIT")
	serviceDiscoveryReq("unregister", "localhost:6667", "DELETE")
	os.Exit(0)
}
