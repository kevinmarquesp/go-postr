package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"

	"github.com/kevinmarquesp/go-postr/internal/router"
)

const Dotenv = ".env"

func main() {
	if err := godotenv.Load(Dotenv); err != nil {
		log.Warn("Could not load the" + Dotenv + " file. Using the system's environment from now on...")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Required port variable was not found in the environment.")
	}

	log.Info("Starting the application at the port " + port + ".")

	if err := router.StartRouter(port); err != nil {
		log.Fatal("Router error.", "error", err)
	}
}
