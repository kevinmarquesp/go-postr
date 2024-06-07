package main

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
	"github.com/kevinmarquesp/go-postr/internal/models"
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

	fmt.Println(db_service)
}
