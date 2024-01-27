package controllers

import (
	"fmt"
	"go-postr/models"
	"log"
	"net/http"
	"strings"
)

const invalidChars string = "~`!@#$%^&*()+={}[]|\\:;\"'<>,.?/"

func wasUsernameAlreadyTaken(username string) (bool, error) {
	rows, err := models.Db.Query("SELECT username FROM public.user WHERE username LIKE $1", username)
	if err != nil {
		log.Println("Couldn't check if this username already exists")
		return false, err
	}
	
	defer rows.Close()

	for rows.Next() {
		var dbUsername string

		if err := rows.Scan(&dbUsername); err != nil {
			log.Println("Couldn't asign username from database to a local variable")
			return false, err

		} else if dbUsername == username {
			return true, nil
		}
	}

	return false, nil
}

func usernameValidationCases(w http.ResponseWriter, wasAlreadyTaken bool, username string) {
	switch {
	case len(username) == 0:
		log.Println("Username Validation :: empty field")
		fmt.Fprintf(w, "")

	case strings.Contains(username, " "):
		log.Println("Username Validation :: space character detected")
		writeFieldValidationResponse(w, "danger", "Space characters aren't allowed")

	case strings.ContainsAny(username, invalidChars):
		log.Println("Username Validation :: invalid characters detected")
		writeFieldValidationResponse(w, "danger", "Use only letters, number and - or _ characters")

	case wasAlreadyTaken:
		log.Println("Username Validation :: that username string was already taken!")
		writeFieldValidationResponse(w, "danger", "Username already taken!")

	default:
		log.Println("Username Validation :: valid user!")
		writeFieldValidationResponse(w, "success", "Valid username")
	}
}

func ValidateUsernameController(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = parseValidationFormFields(w, r); err != nil {
		log.Println("[ERROR]", err)
		return
	}

	username := r.Form.Get("username")
	username = strings.TrimSpace(username)

	wasAlreadyTaken, err := wasUsernameAlreadyTaken(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("[ERROR]", err)

		return
	}

	usernameValidationCases(w, wasAlreadyTaken, username)
}
