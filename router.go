package main

import (
	"go-postr/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer(port string) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", controllers.HomePageController)

	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Println("Could not start the gorilla/mux server")
		log.Panic(err)
	}
}
