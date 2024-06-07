package main

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/kevinmarquesp/go-postr/internal/controllers"
	"github.com/kevinmarquesp/go-postr/internal/models"
)

func StartRouter(db_service models.DatabaseService, port string) error {
	log.Info("Starting the server mux of the application...")

	handler := http.NewServeMux()

	handler.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	pages_controller := controllers.NewPagesController(db_service)

	handler.HandleFunc("GET /", pages_controller.HomePage)

	server := http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	return server.ListenAndServe()
}
