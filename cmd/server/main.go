package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"

	"github.com/kevinmarquesp/go-postr/internal/models"
)

const DOTENV = ".env"

func main() {
	log.Info("Initializing the Go Postr application...")

	if err := godotenv.Load(DOTENV); err != nil {
		log.Warn("Could not load the" + DOTENV + " file, using the system's environment.")
	}

	db := &models.Sqlite{}

	err := db.Connect(os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal("Could not connect to the database.", "err", err)
	}

	log.Print("Hello world!")
}
