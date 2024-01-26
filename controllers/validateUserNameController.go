package controllers

import (
	"log"
	"net/http"
	"strings"
)

const INVALID_CHARS string = "~`!@#$%^&*()+={}[]|\\:;\"'<>,.?/"

type ValidateUsernameResponse struct {
	IsEmpty bool
	Status  string
	Msg     string
}

func ValidateUsernameController(w http.ResponseWriter, r *http.Request) {
	userName := r.URL.Query().Get("user_name")
	userName = strings.TrimSpace(userName)

	switch {
	case len(userName) == 0:
		WriteUsernameValidationStatus(w, "", "")

	case strings.Contains(userName, " "):
		WriteUsernameValidationStatus(w, "danger", "Space characters aren't allowed")	

	case strings.ContainsAny(userName, INVALID_CHARS):
		WriteUsernameValidationStatus(w, "danger", "Use only letters, number and - or _ characters")

	default:
		WriteUsernameValidationStatus(w, "success", "Valid username")
	}
}

func WriteUsernameValidationStatus(w http.ResponseWriter, status, msg string) {
	var isEmpty bool

	if len(status) == 0 || len(msg) == 0 {
		isEmpty = true
	} else {
		isEmpty = false
	}

	err := Tmpl.ExecuteTemplate(w, "AuthStatus", ValidateUsernameResponse{
		IsEmpty: isEmpty,
		Status:  status,
		Msg:     msg,
	})

	if err != nil {
		log.Println("Could not write the AuthStatus component...")
		log.Fatal(err)
	}
}
