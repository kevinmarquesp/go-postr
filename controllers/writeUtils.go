package controllers

import (
	"go-postr/models"
	"log"
	"net/http"
)

func writeFieldValidationResponse(w http.ResponseWriter, bootstrapStatus, infoMessage string) {
	data := models.FieldValidationResponseData{
		BootstrapStatus: bootstrapStatus,
		InfoMessage:     infoMessage,
	}

	log.Printf("Wrinting username validation result: %s, %s\n",
		data.BootstrapStatus, data.InfoMessage)

	err := Tmpl.ExecuteTemplate(w, "Components.FormFieldStatus", data)
	if err != nil {
		log.Println("Could not write the AuthStatus component...")
		log.Fatal(err)
	}
}
