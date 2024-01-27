package controllers

import (
	"errors"
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

func parseValidationFormFields(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return errors.New("Invalid method, expected 'post'")
	}

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	return nil
}

