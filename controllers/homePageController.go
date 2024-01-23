package controllers

import (
	"log"
	"net/http"
)

func HomePageController(w http.ResponseWriter, r *http.Request) {
	err := Tmpl.ExecuteTemplate(w, "HomePage", nil)
	if err != nil {
		log.Println("Could not render the HomePage template, something went wrong")
		log.Fatal(err)
	}
}
