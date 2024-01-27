package controllers

import (
	"log"
	"net/http"
)

func passwordValidationCases(w http.ResponseWriter, length int) {
	switch {
	case length == 0:
		log.Println("Password Validation :: empty field")

	case length < 8:
		log.Println("Password Validation :: too short")
		writeFieldValidationResponse(w, "danger", "You're password is too weak")

	case length >= 8 && length < 12:
		log.Println("Password Validation :: could be better")
		writeFieldValidationResponse(w, "warning", "You could do better than this...")
	
	default:
		log.Println("Password Validation :: password length is ok")
		writeFieldValidationResponse(w, "success", "Eh. Good enough")
	}
}

func ValidatePasswordController(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = parseValidationFormFields(w, r); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("[ERROR]", err)

		return
	}

	password := r.Form.Get("password")
	length := len(password)

	passwordValidationCases(w, length)
}
