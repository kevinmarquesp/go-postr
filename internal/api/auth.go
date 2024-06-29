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
	UNSPECIFIED_AUTHORIZATION_FIELD_ERROR = "a valid username and password is required to authorize"
	INVALID_ACCESS_TOKEN_ERROR            = "the given session token is invalid or was expired"
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

	if err = json.Unmarshal(rawBody, &body); err != nil {
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

	updateSessionTokenWithCredentials := func() error {
		if username == "" || password == "" {
			return errors.New(UNSPECIFIED_AUTHORIZATION_FIELD_ERROR)
		}

		newSessionToken, err := ac.Database.AuthorizeUserWithCredentials(username, password)
		if err != nil {
			utils.WriteGenericJsonError(w, http.StatusUnauthorized, err)
			return nil
		}

		fmt.Fprintf(w, `{ "newSessionToken": "%s" }`, newSessionToken)

		return nil
	}

	// Authorize with acess session token string.

	if sessionToken != "" {
		newSessionToken, err := ac.Database.AuthorizeUserWithSessionToken(username, sessionToken)
		if err != nil {
			if err = updateSessionTokenWithCredentials(); err != nil {
				utils.WriteGenericJsonError(w, http.StatusUnauthorized,
					errors.New(INVALID_ACCESS_TOKEN_ERROR))
			}
			return
		}

		fmt.Fprintf(w, `{ "newSessionToken": "%s" }`, newSessionToken)
		return
	}

	// Authorize with username and password credentials.

	if err = updateSessionTokenWithCredentials(); err != nil {
		utils.WriteGenericJsonError(w, http.StatusUnauthorized, err)
	}
}
