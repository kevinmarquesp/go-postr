package main

import (
	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

const DOTENV = ".env"

func main() {
	log.Info("Initializing the Go Postr application...")

	if err := godotenv.Load(DOTENV); err != nil {
		log.Warn("Could not load the" + DOTENV + " file, using the system's environment.")
	}

	log.Print("Hello world!")
}
