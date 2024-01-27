package controllers

import (
	"fmt"
	"log"
	"net/http"
)

func ValidatePasswordController(w http.ResponseWriter, r *http.Request) {
	password := r.URL.Query().Get("password")
	length := len(password)

	switch {
	case length == 0:
		log.Println("Validating password: empty field")
		fmt.Fprintf(w, "")

	case length < 8:
		log.Println("Validating password: password specified is too short")
		writeFieldValidationResponse(w, "danger", "You're password is too weak")

	case length >= 8 && length < 12:
		log.Println("Validating password: password specifyed could be better")
		writeFieldValidationResponse(w, "warning", "You could do better than this...")
	
	default:
		log.Println("Validating password: password specifyed is ok")
		writeFieldValidationResponse(w, "success", "Eh. Good enough")
	}
}
