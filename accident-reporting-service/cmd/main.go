package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"senpainikolay/pad-sem7/accident-reporting-service/internal/controller"
	"senpainikolay/pad-sem7/accident-reporting-service/internal/models"
	"senpainikolay/pad-sem7/accident-reporting-service/internal/repository"
	"senpainikolay/pad-sem7/accident-reporting-service/internal/service"
	postgres "senpainikolay/pad-sem7/accident-reporting-service/pkg"

	_ "github.com/joho/godotenv/autoload"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"gorm.io/gorm"
)

var db *gorm.DB

var (
	totalRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "pad_req_counter",
			Help: "Count the request made by user.",
		},
		[]string{"Status"},
	)
)

func main() {

	accRepo := repository.NewAccidentRepository(db)
	accService := service.NewAccidentService(accRepo)

	serviceDiscoveryReq("register", os.Getenv("LOCALNAME")+":"+os.Getenv("SERVICE_PORT"), "POST")

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe(":"+os.Getenv("HTTP_PORT"), nil); err != nil {
			log.Fatalf("Failed to start metrics server: %v", err)
		}
	}()

	log.Printf("starting gRPC API server...\n")
	controller.Serve(accService, ":"+os.Getenv("SERVICE_PORT"), totalRequests)

}

func init() {
	db = postgres.NewDBConnection()
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

	prometheus.MustRegister(totalRequests)

}

func serviceDiscoveryReq(route, serviceUrl, method string) {

	url := fmt.Sprintf("http://"+os.Getenv("SERVICE_DISCOVERY_HOST")+":"+os.Getenv("SERVICE_DISCOVERY_PORT")+"/services/%s", route)

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
