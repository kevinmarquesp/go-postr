package main

import (
	"net/http"

	"github.com/kevinmarquesp/go-postr/views/home"
)

func main() {
	app := http.NewServeMux()

	app.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	app.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		home.Home().Render(r.Context(), w)
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: app,
	}

	server.ListenAndServe()
}
