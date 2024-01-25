package controllers

import (
	"log"
	"net/http"
)

func LoginPageController(w http.ResponseWriter, r *http.Request) {
	err := Tmpl.ExecuteTemplate(w, "LoginPage", nil)
	if err != nil {
		log.Println("Could not render the HomePage template, something went wrong")
		log.Fatal(err)
	}
}
