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
	pgHost := os.Getenv("PG_HOST")
	pgPort := os.Getenv("PG_PORT")
	pgUser := os.Getenv("PG_USER")
	pgPassword := os.Getenv("PG_PASSWORD")
	pgDbname := os.Getenv("PG_DBNAME")

	models.OpenGlobalConnection(pgHost, pgPort, pgUser, pgPassword, pgDbname)

	StartServer(":" + port)
}
