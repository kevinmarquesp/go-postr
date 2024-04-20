package main

import (
	"go-postr/db"
	"go-postr/router"
	"go-postr/utils"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(db.Dotenv)
	if err != nil {
		log.Println("The " + dotenv + " file was not found, it will use the system's environment then.")
	}

	creds, err := db.DefaultCredentials()
	if err != nil {
		log.Fatalln("Could not retrieve postgres environment variables.")
	}

	// This function sets up a global variable inside this package. All database
	// interactions will trhow an error if this function was not ran previously.

	err = db.Connect(creds)
	if err != nil {
		log.Fatalln("Postgres connection failed.")
	}

	port, err := utils.RequireEnv("PORT")
	if err != nil {
		log.Fatalln("PORT environment variable not found.")
	}

	log.Println("Starting up the server router.")
	router.InitRouter(":" + port) // This ":" is required by the http.Server{} struct.
}
