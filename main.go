package main

import (
	"errors"
	"go-postr/db"
	"go-postr/router"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const dotenv = ".env"

func main() {
	err := godotenv.Load(dotenv)
	if err != nil {
		log.Println("WARNING: Could not find the", dotenv, "file")
		log.Println("WARNING: This code will use the system's environment")
	}

	port, err := requireEnv("PORT")
	pgHost, err := requireEnv("POSTGRES_HOST")
	pgPort, err := requireEnv("POSTGRES_PORT")
	pgUsernmae, err := requireEnv("POSTGRES_USER")
	pgPassword, err := requireEnv("POSTGRES_PASSWORD")
	pgDatabase, err := requireEnv("POSTGRES_DB")

	if err != nil {
		log.Println("ERROR: Required variables not satisfied")
		log.Fatalln(err)
	}

	port = ":" + port //add a little ":" to be compatible with the http.ListenAndServe() function

	_ = db.Connect(db.ConnCredentials{
		Host:         pgHost,
		Port:         pgPort,
		Username:     pgUsernmae,
		Password:     pgPassword,
		DatabaseName: pgDatabase,
	})

	log.Println("Starting the server router")
	router.InitRouter(port)
}
