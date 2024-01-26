package utils

import (
	"go-postr/models"
	"html/template"
	"log"
	"net/http"
)

func WriteUsernameValidationStatus(w http.ResponseWriter, tmpl *template.Template, data models.ValidateUsernameResponse) {
	log.Printf("Wrinting username validation result: %t, %s, %s\n",
		data.IsEmpty, data.BootstrapStatus, data.InfoMessage)

	err := tmpl.ExecuteTemplate(w, "AuthStatus", data)
	if err != nil {
		log.Println("Could not write the AuthStatus component...")
		log.Fatal(err)
	}
}
