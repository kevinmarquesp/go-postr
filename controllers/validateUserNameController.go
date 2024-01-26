package controllers

import (
	"log"
	"net/http"
	"strings"
)

const INVALID_CHARS string = "~`!@#$%^&*()+={}[]|\\:;\"'<>,.?/"

type responseContent struct {
	IsEmpty bool
	Status  string
	Msg     string
}

func ValidateUserNameController(w http.ResponseWriter, r *http.Request) {
	userName := r.URL.Query().Get("user_name")
	userName = strings.TrimSpace(userName)

	switch {
	case len(userName) == 0:
		writeStatusComponent(w, "", "")

	case strings.Contains(userName, " "):
		writeStatusComponent(w, "danger", "Space characters aren't allowed")	

	case strings.ContainsAny(userName, INVALID_CHARS):
		writeStatusComponent(w, "danger", "Use only letters, number and - or _ characters")

	default:
		writeStatusComponent(w, "success", "Valid username")
	}
}

func writeStatusComponent(w http.ResponseWriter, status, msg string) {
	var isEmpty bool

	if len(status) == 0 || len(msg) == 0 {
		isEmpty = true
	} else {
		isEmpty = false
	}

	err := Tmpl.ExecuteTemplate(w, "AuthStatus", responseContent{
		IsEmpty: isEmpty,
		Status:  status,
		Msg:     msg,
	})

	if err != nil {
		log.Println("Could not write the AuthStatus component...")
		log.Fatal(err)
	}
}
