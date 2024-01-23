package main

import (
	"go-postr/models"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Could not read `.env` file")
		log.Panic(err)
	}

	port := os.Getenv("PORT")

	conn := models.Connection{
		Host:     os.Getenv("PG_HOST"),
		Port:     os.Getenv("PG_PORT"),
		User:     os.Getenv("PG_USER"),
		Password: os.Getenv("PG_PASSWORD"),
		Dbname:   os.Getenv("PG_DBNAME"),
	}

	models.OpenGlobalConnection(conn)
	models.ExecuteMigration()

	StartServer(":" + port)
}
