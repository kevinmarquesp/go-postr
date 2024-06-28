package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/kevinmarquesp/go-postr/internal/data"
	"github.com/kevinmarquesp/go-postr/internal/models"
	"github.com/kevinmarquesp/go-postr/internal/utils"
)

type AuthController struct {
	Database models.DatabaseProvider
}

func (ac AuthController) RegisterNewUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	rawBody, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteJsonError(w, http.StatusInternalServerError, err)

		return
	}

	var body data.RegisterCredentialsIncome

	err = json.Unmarshal(rawBody, &body)
	if err != nil {
		utils.WriteJsonError(w, http.StatusInternalServerError, err)

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
