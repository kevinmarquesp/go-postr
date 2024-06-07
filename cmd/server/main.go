package main

import (
	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
	"github.com/kevinmarquesp/go-postr/internal/models"
	"github.com/kevinmarquesp/go-postr/internal/utils"
)

const DOTENV = ".env"

func main() {
	log.Info("Application initialized.")

	if err := godotenv.Load(DOTENV); err != nil {
		log.Error("Could not load the "+DOTENV+" file.", "error", err)
	}

	db_service := &models.Postgres{}

	if err := db_service.Connect(); err != nil {
		log.Fatal("Could not connect to the specifyed database service.", "error", err)
	}

	port, err := utils.ProtectedGetenv("PORT")
	if err != nil {
		log.Fatal("Could not locate the server's port variable on the system.", "error", err)
	}

	if err := StartRouter(db_service, port); err != nil {
		log.Fatal("Unexpected error, server shut down.", "error", err)
	}
}
