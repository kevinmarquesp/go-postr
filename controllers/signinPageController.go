package controllers

import (
	"log"
	"net/http"
)

func SigninPageController(w http.ResponseWriter, r *http.Request) {
	err := Tmpl.ExecuteTemplate(w, "Signin", nil)
	if err != nil {
		log.Fatal("[ERROR]", err)
	}
}
