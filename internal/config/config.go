package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AdAstraPort		string
	NasaURL      	string
	NasaAPIKey   	string
	SpaceXURL    	string
	SpaceXAPIKey 	string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, loading from environment")
	}

	cfg := &Config{
		AdAstraPort: 	os.Getenv("ADASTRA_PORT"),
		NasaURL: 		os.Getenv("NASA_URL"),
		NasaAPIKey: 	os.Getenv("NASA_APIKEY"),
		SpaceXURL: 		os.Getenv("SPACEX_URL"),
		SpaceXAPIKey: 	os.Getenv("SPACEX_APIKEY"),
	}

	return cfg, nil
}