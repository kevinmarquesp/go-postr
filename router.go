package main

import (
	"go-postr/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer(port string) {
	router := mux.NewRouter().StrictSlash(true)

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	router.HandleFunc("/", controllers.HomePageController)
	router.HandleFunc("/login", controllers.LoginPageController)
	
	router.HandleFunc("/auth/validate/user", controllers.ValidateUserNameController)

	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Println("Could not start the gorilla/mux server")
		log.Panic(err)
	}
}
