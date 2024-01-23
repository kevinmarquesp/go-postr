package main

import (
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

	StartServer(":" + port)
}
