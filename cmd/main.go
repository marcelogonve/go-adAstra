package main

import (
	"log"

	"go-adAstra/internal/config"
	"go-adAstra/internal/routes"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	router := routes.SetupRouter()

	port := ":" + cfg.AdAstraPort

	router.Run(port)
}
