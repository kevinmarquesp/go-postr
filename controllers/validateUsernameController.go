package controllers

import (
	"fmt"
	"go-postr/models"
	"log"
	"net/http"
	"strings"
)

const INVALID_CHARS string = "~`!@#$%^&*()+={}[]|\\:;\"'<>,.?/"

func isUsernameAlreadyTaken(username string) (bool, error) {
	rows, err := models.Db.Query("SELECT user_name FROM public.user WHERE user_name LIKE $1", username)
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

func ValidateUsernameController(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = parseValidationFormFields(w, r); err != nil {
		log.Println("[ERROR]", err)
		return
	}

	username := r.Form.Get("username")
	username = strings.TrimSpace(username)

	wasAlreadyTaken, err := isUsernameAlreadyTaken(username)
	if err != nil {
		writeFieldValidationResponse(w, "warning", "Something went wrong at the server")
		log.Println(err)
		return
	}

	switch {
	case len(username) == 0:
		log.Println("Validating username: empty field")
		fmt.Fprintf(w, "")

	case strings.Contains(username, " "):
		log.Println("Validating username: space character detected")
		writeFieldValidationResponse(w, "danger", "Space characters aren't allowed")

	case strings.ContainsAny(username, INVALID_CHARS):
		log.Println("Validating username: invalid characters detected")
		writeFieldValidationResponse(w, "danger", "Use only letters, number and - or _ characters")

	case wasAlreadyTaken:
		log.Println("Validating username: that username string was already taken!")
		writeFieldValidationResponse(w, "danger", "Username already taken!")

	default:
		log.Println("Validating username: valid user!")
		writeFieldValidationResponse(w, "success", "Valid username")
	}
}
