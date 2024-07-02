package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
	"github.com/kevinmarquesp/go-postr/internal/models"
)

// NOTE: All functions here should handle the errors by their own!

const DOTENV = ".env"

func main() {
	db := DatabaseConnect()

	InsertDummyUsers(db)
}

func DatabaseConnect() models.GenericDatabaseProvider {
	if err := godotenv.Load(DOTENV); err != nil {
		log.Warn("Could not load the" + DOTENV + " file, using the system's environment.")
	}

	db := &models.Sqlite{}

	if err := db.Connect(os.Getenv("DATABASE_URL")); err != nil {
		log.Fatal("Could not connect to the database.", "err", err)
	}

	return db
}
