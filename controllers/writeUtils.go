package controllers

import (
	"fmt"
	"go-postr/models"
	"log"
	"net/http"
)

func writeAnEmptyString(w http.ResponseWriter) {
	fmt.Fprintf(w, "")
}

func writeFieldValidationResponse(w http.ResponseWriter, data models.FieldValidationResponseData) {
	log.Printf("Wrinting username validation result: %t, %s, %s\n",
		data.IsEmpty, data.BootstrapStatus, data.InfoMessage)

	err := Tmpl.ExecuteTemplate(w, "AuthStatus", data)
	if err != nil {
		log.Println("Could not write the AuthStatus component...")
		log.Fatal(err)
	}
}

func writeFieldValidationResponseWraper(w http.ResponseWriter, bootstrapStatus, infoMessage string) {
	writeFieldValidationResponse(w, models.FieldValidationResponseData{
		BootstrapStatus: bootstrapStatus,
		InfoMessage:     infoMessage,
	})
}
