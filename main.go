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

		rows, err := db.Connection().Query(`SELECT username FROM "user" WHERE username LIKE $1`,
			"%" + query + "%")
		if err != nil {
			log.Error("Could not fetch usernames like " + query, "error", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

			return
		}

		list := ""

		for rows.Next() {
			var username string

			err = rows.Scan(&username)
			if err != nil {
				log.Error("Could not scan db column...", "error", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

				return
			}

			list += fmt.Sprintf(`<li><a href="/u/%s">%s</a></li>`, username, username)
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

	port := ":" + os.Getenv("PORT")

	if port == ":" {
		log.Warn("PORT variable not specified, using default", "port", ":8000")
		port = ":8000"
	}

	log.Info("Starting the server router...")
	router(port)
}
