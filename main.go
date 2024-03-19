package main

import (
	"go-postr/templates"
	"net/http"
	"os"
)


func main() {
	port := ":" + os.Getenv("PORT")

	if port == ":" {
		port = ":8000"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ := templates.NewTemplateRenderer()

		templ.Render(w, "Index", nil)
	})

	http.ListenAndServe(port, nil)
}
