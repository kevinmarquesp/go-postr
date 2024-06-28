package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/kevinmarquesp/go-postr/internal/models"
)

// TODO: Move this struct to a JSON data pakcage.

type RegisterCredentials struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthController struct {
	Database models.DatabaseProvider
}

func (ac AuthController) RegisterNewUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	rawBody, err := io.ReadAll(r.Body)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)

		return
	}

	var body RegisterCredentials

	err = json.Unmarshal(rawBody, &body)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)

		return
	}

	// TODO: Create a functino to validate each input field.

	log.Print(body)

	// TODO: Check if the user exists on the database.
	// TODO: Generate a new access token for that user.
	// TODO: Register that token in the database.
	// TODO: Return the token string to the final user.

	fmt.Fprint(w, `{ "message": "Registering a new user" }`)
}

// TODO: Move this function to an utils package.

func WriteError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	fmt.Fprintf(w, `{ "error": "%s" }`, err)
}
