package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/kevinmarquesp/go-postr/internal/models"
	"github.com/kevinmarquesp/go-postr/internal/utils"
)

const (
	INVALID_ACCESS_TOKEN_ERROR            = "the given session token is invalid or was expired"
	UNSPECIFIED_AUTHORIZATION_FIELD_ERROR = "a valid username and password is required to authorize"
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

	var form models.RegisterForm

	if err = json.Unmarshal(rawBody, &form); err != nil {
		utils.WriteGenericJsonError(w, http.StatusInternalServerError, err)
		return
	}

	form.Fullname = strings.Trim(form.Fullname, " ")
	form.Username = strings.Trim(form.Username, " ")
	form.Password = strings.Trim(form.Password, " ")

	// Register and respond.

	response, err := ac.Database.RegisterNewUser(form)
	if err != nil {
		utils.WriteGenericJsonError(w, http.StatusBadRequest, err)
		return
	}

	responseJson, err := utils.JsonMarshalString(response)
	if err != nil {
		utils.WriteGenericJsonError(w, http.StatusConflict, err)
		return
	}

	fmt.Fprint(w, string(responseJson))
}

func (ac AuthController) RefreshUserSessionToken(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	rawBody, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteGenericJsonError(w, http.StatusInternalServerError, err)
		return
	}

	var body models.Auth

	err = json.Unmarshal(rawBody, &body)
	if err != nil {
		utils.WriteGenericJsonError(w, http.StatusInternalServerError, err)
		return
	}

	username := strings.Trim(body.Username, " ")
	password := strings.Trim(body.Password, " ")
	sessionToken := strings.Trim(body.SessionToken, " ")

	// Authorize with acess session token string.

	if sessionToken != "" {
		newSessionToken, err := ac.Database.AuthorizeUserWithSessionToken(sessionToken)
		if err != nil {
			if err = ac.updateSessionTokenWithCredentials(w, username, password); err != nil {
				utils.WriteGenericJsonError(w, http.StatusUnauthorized,
					errors.New(INVALID_ACCESS_TOKEN_ERROR))
			}
			return
		}

		response := models.SessionToken{
			SessionToken: newSessionToken,
		}

		responseJsonBytes, err := json.Marshal(response)
		if err != nil {
			utils.WriteGenericJsonError(w, http.StatusInternalServerError, err)
		}

		fmt.Fprint(w, string(responseJsonBytes))
		return
	}

	// Authorize with username and password credentials.

	if err = ac.updateSessionTokenWithCredentials(w, username, password); err != nil {
		utils.WriteGenericJsonError(w, http.StatusUnauthorized, err)
	}
}

func (ac AuthController) updateSessionTokenWithCredentials(w http.ResponseWriter, username, password string) error {
	if username == "" || password == "" {
		return errors.New(UNSPECIFIED_AUTHORIZATION_FIELD_ERROR)
	}

	newSessionToken, err := ac.Database.AuthorizeUserWithCredentials(username, password)
	if err != nil {
		utils.WriteGenericJsonError(w, http.StatusUnauthorized, err)
		return nil
	}

	response := models.SessionToken{
		SessionToken: newSessionToken,
	}

	responseJsonBytes, err := json.Marshal(response)
	if err != nil {
		return err
	}

	fmt.Fprint(w, string(responseJsonBytes))

	return nil
}
