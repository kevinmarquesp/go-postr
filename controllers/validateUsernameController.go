package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const INVALID_CHARS string = "~`!@#$%^&*()+={}[]|\\:;\"'<>,.?/"

func ValidateUsernameController(w http.ResponseWriter, r *http.Request) {
	userName := r.URL.Query().Get("user_name")
	userName = strings.TrimSpace(userName)

	switch {
	case len(userName) == 0:
		log.Println("Validating username: empty field")
		fmt.Fprintf(w, "")

	case strings.Contains(userName, " "):
		log.Println("Validating username: space character detected")
		writeFieldValidationResponse(w, "danger", "Space characters aren't allowed")

	case strings.ContainsAny(userName, INVALID_CHARS):
		log.Println("Validating username: invalid characters detected")
		writeFieldValidationResponse(w, "danger", "Use only letters, number and - or _ characters")

	default:
		log.Println("Validating username: valid user!")
		writeFieldValidationResponse(w, "success", "Valid username")
	}
}
