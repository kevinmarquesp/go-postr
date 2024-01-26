package controllers

import (
	"log"
	"net/http"
	"strings"
)

const INVALID_CHARS string = "~`!@#$%^&*()+={}[]|\\:;\"'<>,.?/"

type ValidateUsernameResponse struct {
	IsEmpty         bool
	BootstrapStatus string
	InfoMessage     string
}

func ValidateUsernameController(w http.ResponseWriter, r *http.Request) {
	userName := r.URL.Query().Get("user_name")
	userName = strings.TrimSpace(userName)

	switch {
	case len(userName) == 0:
		log.Println("Validating username: empty field")

		WriteUsernameValidationStatus(w, ValidateUsernameResponse{
			IsEmpty: true,
		})

	case strings.Contains(userName, " "):
		log.Println("Validating username: space character detected")

		WriteUsernameValidationStatus(w, ValidateUsernameResponse{
			BootstrapStatus: "danger",
			InfoMessage:     "Space characters aren't allowed",
		})

	case strings.ContainsAny(userName, INVALID_CHARS):
		log.Println("Validating username: invalid characters detected")

		WriteUsernameValidationStatus(w, ValidateUsernameResponse{
			BootstrapStatus: "danger", 
			InfoMessage:     "Use only letters, number and - or _ characters",
		})

	default:
		log.Println("Validating username: valid user!")

		WriteUsernameValidationStatus(w, ValidateUsernameResponse{
			BootstrapStatus: "success", 
			InfoMessage:     "Valid username",
		})
	}
}

func WriteUsernameValidationStatus(w http.ResponseWriter, data ValidateUsernameResponse) {
	log.Printf("Wrinting username validation result: %t, %s, %s\n",
		data.IsEmpty, data.BootstrapStatus, data.InfoMessage)

	err := Tmpl.ExecuteTemplate(w, "AuthStatus", data)
	if err != nil {
		log.Println("Could not write the AuthStatus component...")
		log.Fatal(err)
	}
}
