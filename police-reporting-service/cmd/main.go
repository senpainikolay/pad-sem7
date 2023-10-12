package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"senpainikolay/pad-sem7/police-reporting-service/internal/controller"
	"senpainikolay/pad-sem7/police-reporting-service/internal/repository"
	"senpainikolay/pad-sem7/police-reporting-service/internal/service"
	mongodbp "senpainikolay/pad-sem7/police-reporting-service/pkg"
	"syscall"
)

func main() {
	dbName := "test_db"
	MONGO_CLIENT := mongodbp.NewDBConnection()

	policeRepo := repository.NewPoliceRepository(MONGO_CLIENT, dbName)
	policeService := service.NewPoliceService(policeRepo)

	serviceDiscoveryReq("register", "localhost:6666", "POST")

	// in case of error
	defer serviceDiscoveryReq("unregister", "localhost:6666", "DELETE")

	// in case of force stop
	go signalUnregisterThread()

	controller.Serve(policeService, ":6666")

}

func serviceDiscoveryReq(route, serviceUrl, method string) {

	url := fmt.Sprintf("http://localhost:8000/services/%s", route)

	requestData := map[string]string{
		"service_type": "police-reporting",
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
	serviceDiscoveryReq("unregister", "localhost:6666", "DELETE")
	os.Exit(0)
}
