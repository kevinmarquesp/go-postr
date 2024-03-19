package main

import (
	"errors"
	"fmt"
	"go-postr/db"
	"go-postr/templates"
	"net/http"
	"os"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

const dotenv = ".env"

func router(port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ := templates.NewTemplateRenderer()

		templ.Render(w, "Index", nil)
	})

	http.HandleFunc("/search/user", func(w http.ResponseWriter, r *http.Request) {
		v := r.URL.Query()
		query := v.Get("query")

		if len(query) == 0 {
			fmt.Fprintf(w, "")  // insert an empty string in the results tag

			return
		}

		list, err := db.SearchByUsername(query)
		if err != nil {
			log.Error("Couldn't search for user " + query, "error", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

			return
		}

		fmt.Fprintf(w, list)
	})

	log.Info("Listening", "url", "http://localhost" + port)
	http.ListenAndServe(port, nil)
}

func requireEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", errors.New("required environment variable " + key + " doesn't exist or is empty")
	}

	return value, nil
}

func main() {
	err := godotenv.Load(dotenv)
	if err != nil {
		log.Error("Couldn't load the " + dotenv + " file", "error", err)
		log.Warn("The server will use the system's environment!")
	}

	port, err := requireEnv("PORT")
	pgHost, err := requireEnv("POSTGRES_HOST")
	pgPort, err := requireEnv("POSTGRES_PORT")
	pgUsernmae, err := requireEnv("POSTGRES_USER")
	pgPassword, err := requireEnv("POSTGRES_PASSWORD")
	pgDatabase, err := requireEnv("POSTGRES_DB")

	if err != nil {
		log.Fatal("Required variables not satisfied", "error", err)
		os.Exit(1)
	}

	port = ":" + port  //add a little ":" to be compatible with the http.ListenAndServe() function

	_ = db.Connect(db.ConnCredentials{
		Host: pgHost,
		Port: pgPort,
		Username: pgUsernmae,
		Password: pgPassword,
		DatabaseName: pgDatabase,
	})

	log.Info("Starting the server router...")
	router(port)
}
