package main

import (
	"go-postr/templates"
	"net/http"
	"os"

	"github.com/charmbracelet/log"
)

func router(port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ := templates.NewTemplateRenderer()

		templ.Render(w, "Index", nil)
	})

	http.HandleFunc("/search/user", func(w http.ResponseWriter, r *http.Request) {
		v := r.URL.Query()

		log.Info(v.Get("query"))
	})

	log.Info("Listening", "url", "http://localhost" + port)
	http.ListenAndServe(port, nil)
}

func main() {
	port := ":" + os.Getenv("PORT")

	if port == ":" {
		log.Warn("PORT variable not specified, using default", "port", ":8000")
		port = ":8000"
	}

	log.Info("Starting the server router...")
	router(port)
}
