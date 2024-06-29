package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/kevinmarquesp/go-postr/internal/data"
	"github.com/kevinmarquesp/go-postr/internal/models"
	"github.com/kevinmarquesp/go-postr/internal/utils"
)

const (
	UNSPECIFIED_AUTHORIZATION_FIELD_ERROR = "username and password or a session token is required to authorize"
)

type AuthController struct {
	Database models.GenericDatabaseProvider
}

func (ac AuthController) RegisterNewUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	rawBody, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteGenericJsonError(w, http.StatusInternalServerError, err)
		return
	}

	var body data.RegisterNewUserBodyCredentialsBody

	err = json.Unmarshal(rawBody, &body)
	if err != nil {
		utils.WriteGenericJsonError(w, http.StatusInternalServerError, err)
		return
	}

	// Field validation.

	username := strings.Trim(body.Username, " ")
	password := strings.Trim(body.Password, " ")

	if err = utils.ValidateUsernameString(username); err != nil {
		utils.WriteGenericJsonError(w, http.StatusBadRequest, err)
		return
	}

	if err = utils.ValidatePasswordString(password); err != nil {
		utils.WriteGenericJsonError(w, http.StatusBadRequest, err)
		return
	}

	// Register and respond.

	sessionToken, err := ac.Database.RegisterNewUser(username, password)
	if err != nil {
		utils.WriteGenericJsonError(w, http.StatusBadRequest, err)
		return
	}

	successfulReponseData := data.RegisterNewUserSuccessfulResponse{
		Username:     username,
		SessionToken: sessionToken,
	}

	successfulReponseJsonData, err := json.Marshal(successfulReponseData)
	if err != nil {
		utils.WriteGenericJsonError(w, http.StatusConflict, err)
		return
	}

	fmt.Fprint(w, string(successfulReponseJsonData))
}

func (ac AuthController) UpdateUserSessionToken(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	rawBody, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteGenericJsonError(w, http.StatusInternalServerError, err)
		return
	}

	var body data.UpdateUserSessionTokenCredentialsBody

	err = json.Unmarshal(rawBody, &body)
	if err != nil {
		utils.WriteGenericJsonError(w, http.StatusInternalServerError, err)
		return
	}

	username := strings.Trim(body.Username, " ")
	password := strings.Trim(body.Password, " ")
	sessionToken := strings.Trim(body.SessionToken, " ")

	// Authorize with acess session token string.

	if username == "" && password == "" && len(sessionToken) > 0 {
		newSessionToken, err := ac.Database.AuthorizeUserWithSessionToken(username, sessionToken)
		if err != nil {
			utils.WriteGenericJsonError(w, http.StatusUnauthorized, err)
			return
		}

		fmt.Fprintf(w, `{ "newSessionToken": "%s" }`, newSessionToken)
		return
	}

	if (username == "" && password != "") || (username != "" && password == "") {
		utils.WriteGenericJsonError(w, http.StatusUnauthorized,
			errors.New(UNSPECIFIED_AUTHORIZATION_FIELD_ERROR))
		return
	}

	// TODO: Authorize with user credentials.
}
