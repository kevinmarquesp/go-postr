package main

import (
	"net/http"
	"go-postr/templates"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ := templates.NewTemplateRenderer()

		templ.Render(w, "Index", nil)
	})

	http.ListenAndServe(":8000", nil)
}
