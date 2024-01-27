package controllers

import (
	"fmt"
	"log"
	"net/http"
)

func ValidatePasswordController(w http.ResponseWriter, r *http.Request) {
	password := r.URL.Query().Get("password")

	switch {
	case len(password) < 8:
		log.Println("Validating password: password specified is too short")
		writeFieldValidationResponse(w, "danger", "You're password is too weak")

	case len(password) >= 8 && len(password) < 12:
		log.Println("Validating password: password specifyed could be better")
		writeFieldValidationResponse(w, "warning", "You could do better than this...")

	case len(password) >= 12:
		log.Println("Validating password: password specifyed is ok")
		writeFieldValidationResponse(w, "success", "Eh. Good enough")
	
	default:
		log.Println("Validating password: empty field")
		fmt.Fprintf(w, "")
	}
}
