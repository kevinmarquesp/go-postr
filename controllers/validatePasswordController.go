package controllers

import (
	"log"
	"net/http"
)

func ValidatePasswordController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeFieldValidationResponse(w, "warning", "Server error: nvalid request method, expected POST")
		log.Println("Invalid request method, expected POST")

		return
	}

	err := r.ParseForm()
	if err != nil {
		writeFieldValidationResponse(w, "warning", "Server error: could not parse form values")

		log.Println("Could not parse the form values")
		log.Println(err)

		return
	}

	password := r.Form.Get("password")
	length := len(password)

	switch {
	case length == 0:
		log.Println("Validating password: empty field")

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
