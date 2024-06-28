package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

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

	// Field validation.

	username := strings.Trim(body.Username, " ")
	password := strings.Trim(body.Password, " ")

	if err = utils.ValidateUsernameString(username); err != nil {
		utils.WriteJsonError(w, http.StatusBadRequest, err)

		return
	}

	if err = utils.ValidatePasswordString(password); err != nil {
		utils.WriteJsonError(w, http.StatusBadRequest, err)

		return
	}

	// Register and respond.

	sessionToken, err := ac.Database.RegisterNewUser(username, password)
	if err != nil {
		utils.WriteJsonError(w, http.StatusInternalServerError, err)

		return
	}

	successfulReponseData := data.RegisterSuccessfulSessionTokenResponse{
		Username:     username,
		SessionToken: sessionToken,
	}

	successfulReponseJsonData, err := json.Marshal(successfulReponseData)
	if err != nil {
		utils.WriteJsonError(w, http.StatusConflict, err)

		return
	}

	fmt.Fprint(w, string(successfulReponseJsonData))
}
