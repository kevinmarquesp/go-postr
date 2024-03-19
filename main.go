package main

import (
	"go-postr/templates"
	"net/http"
	"os"
)

func router(port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ := templates.NewTemplateRenderer()

		templ.Render(w, "Index", nil)
	})

	http.ListenAndServe(port, nil)
}

func main() {
	port := ":" + os.Getenv("PORT")

	if port == ":" {
		port = ":8000"
	}

	router(port)
}
