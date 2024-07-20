package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/kevinmarquesp/go-postr/internal/models"
	"github.com/kevinmarquesp/go-postr/internal/services"
)

type ApiUserHandler struct {
	AuthenticationService services.AuthenticationService
}

func NewApiUserHandler(authenticationService services.AuthenticationService) ApiUserHandler {
	return ApiUserHandler{
		AuthenticationService: authenticationService,
	}
}

func (au ApiUserHandler) RegisterNewUserWithCredentials(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var props struct {
		Name         string `json:"name"`
		Username     string `json:"username"`
		Email        string `json:"email"`
		Password     string `json:"password"`
		Confirmation string `json:"confirmation"`
	}

	err := json.NewDecoder(r.Body).Decode(&props)
	if err != nil {
		models.WriteHttpJsonError(w, http.StatusBadRequest, err,
			"controllers.api-user", "The request body is invalid.")
		return
	}

	createdUser, err := au.AuthenticationService.AuthenticateWithCredentials(
		strings.Trim(props.Name, " "),
		strings.Trim(props.Username, " "),
		strings.Trim(props.Email, " "),
		strings.Trim(props.Password, " "),
		strings.Trim(props.Confirmation, " "),
	)
	if err != nil {
		models.WriteHttpJsonError(w, http.StatusInternalServerError, err,
			"services.authentication", "Could not register the request user to the database.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdUser)
}

func (au ApiUserHandler) FetchUserDataByUsername(w http.ResponseWriter, r *http.Request) {
}

func (au ApiUserHandler) UpdateProfileDetails(w http.ResponseWriter, r *http.Request) {
}

func (au ApiUserHandler) RemoveRegisteredUser(w http.ResponseWriter, r *http.Request) {
}