package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Config holds the application configuration
type Config struct {
	DBUser     string `yaml:"db_user"`
	DBPassword string `yaml:"db_password"`
	DBHost     string `yaml:"db_host"`
	DBName     string `yaml:"db_name"`
}

var AppConfig Config

// LoadConfig loads the configuration from config.yaml
func LoadConfig() {
	file, err := os.Open("config/config.yaml")
	if err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&AppConfig); err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}
}
