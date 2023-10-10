package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	UserDB         string `yaml:"POSTGRES_USER"`
	PostgresPW     string `yaml:"POSTGRES_PASSWORD"`
	PostgresHost   string `yaml:"POSTGRES_HOST"`
	PostgresPort   string `yaml:"POSTGRES_PORT"`
	PostgresDBName string `yaml:"POSTGRES_DB_NAME"`
	ServicePort    string `yaml:"SERVICE_PORT"`
}

func NewConfig(yamlFilePath string) Config {

	data, err := os.ReadFile(yamlFilePath)
	if err != nil {
		log.Fatal(err)
	}
	var config Config

	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatal(err)
	}
	return config
}
