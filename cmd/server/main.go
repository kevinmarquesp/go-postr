package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"

	"github.com/kevinmarquesp/go-postr/internal/models"
	"github.com/kevinmarquesp/go-postr/internal/router"
)

const DOTENV = ".env"

func main() {
	log.Info("Initializing the Go Postr application...")

	if err := godotenv.Load(DOTENV); err != nil {
		log.Warn("Could not load the" + DOTENV + " file, using the system's environment.")
	}

	db := &models.Sqlite{}

	if err := db.Connect(os.Getenv("DATABASE_URL")); err != nil {
		log.Fatal("Could not connect to the database.", "err", err)
	}

	if err := router.InitRouter(os.Getenv("PORT"), db); err != nil {
		log.Fatal("Router error.", "err", err)
	}
}
