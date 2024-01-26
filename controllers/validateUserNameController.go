package controllers

import (
	"fmt"
	"go-postr/models"
	"go-postr/utils"
	"log"
	"net/http"
	"strings"
)

const INVALID_CHARS string = "~`!@#$%^&*()+={}[]|\\:;\"'<>,.?/"

func writeAnEmptyString(w http.ResponseWriter) {
	fmt.Fprintf(w, "")
}

func writeStatusMessage(w http.ResponseWriter, bootstrapStatus, infoMessage string) {
	utils.WriteUsernameValidationStatus(w, Tmpl, models.ValidateUsernameResponse{
		BootstrapStatus: bootstrapStatus,
		InfoMessage:     infoMessage,
	})
}

func ValidateUsernameController(w http.ResponseWriter, r *http.Request) {
	userName := r.URL.Query().Get("user_name")
	userName = strings.TrimSpace(userName)

	switch {
	case len(userName) == 0:
		log.Println("Validating username: empty field")
		writeAnEmptyString(w)

	case strings.Contains(userName, " "):
		log.Println("Validating username: space character detected")
		writeStatusMessage(w, "danger", "Space characters aren't allowed")

	case strings.ContainsAny(userName, INVALID_CHARS):
		log.Println("Validating username: invalid characters detected")
		writeStatusMessage(w, "danger", "Use only letters, number and - or _ characters")

	default:
		log.Println("Validating username: valid user!")
		writeStatusMessage(w, "success", "Valid username")
	}
}
