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
	FULLNAME_REGITER_FIELD_NOT_SPECIFIED_ERROR = "a fullname should be provided to register the user"
	UNSPECIFIED_AUTHORIZATION_FIELD_ERROR      = "a valid username and password is required to authorize"
	INVALID_ACCESS_TOKEN_ERROR                 = "the given session token is invalid or was expired"
)

// AuthController handles authentication-related HTTP requests.
//
// This controller uses a GenericDatabaseProvider to perform operations
// such as registering new users and managing user sessions.
//
// Fields:
// - Database: An instance of a type that implements the models.GenericDatabaseProvider interface.
//
// Example usage:
//
//	func main() {
//	    db := &models.Sqlite{}
//	    authController := AuthController{Database: db}
//
//	    db.Connect("database_url")
//
//	    http.HandleFunc("/register", authController.RegisterNewUser)
//	    http.ListenAndServe(":8080", nil)
//	}
type AuthController struct {
	Database models.GenericDatabaseProvider
}

// This method reads the request body to obtain the user's registration details,
// validates the input, registers the new user in the database, and responds with
// the registration details including a session token.
//
// Example request body (JSON):
//
//	{
//	    "fullname": "John Doe",
//	    "username": "johndoe",
//	    "password": "Password123!"
//	}
//
// Example response body (JSON):
//
//	{
//	    "username": "johndoe",
//	    "public_id": "some-unique-public-id",
//	    "session_token": "some-session-token"
//	}
//
// Possible error responses:
//   - 400 Bad Request: If any of the input validation checks fail.
//   - 409 Conflict: If there is an error while generating the response JSON.
//   - 500 Internal Server Error: If there is an error reading the request body or unmarshaling JSON.
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

	fullname := strings.Trim(body.Fullname, " ")
	username := strings.Trim(body.Username, " ")
	password := strings.Trim(body.Password, " ")

	if len(fullname) == 0 {
		utils.WriteGenericJsonError(w, http.StatusBadRequest,
			errors.New(FULLNAME_REGITER_FIELD_NOT_SPECIFIED_ERROR))
		return
	}

	if err = utils.ValidateUsernameString(username); err != nil {
		utils.WriteGenericJsonError(w, http.StatusBadRequest, err)
		return
	}

	if err = utils.ValidatePasswordString(password); err != nil {
		utils.WriteGenericJsonError(w, http.StatusBadRequest, err)
		return
	}

	// Register and respond.

	publicID, sessionToken, err := ac.Database.RegisterNewUser(fullname, username, password)
	if err != nil {
		utils.WriteGenericJsonError(w, http.StatusBadRequest, err)
		return
	}

	response := data.RegisterNewUserSuccessfulResponse{
		Username:     username,
		PublicID:     publicID,
		SessionToken: sessionToken,
	}

	responseJsonBytes, err := json.Marshal(response)
	if err != nil {
		utils.WriteGenericJsonError(w, http.StatusConflict, err)
		return
	}

	fmt.Fprint(w, string(responseJsonBytes))
}

// This method reads the request body to obtain the user's credentials and session token,
// validates the input, authorizes the user, and responds with a new refreshed session token.
//
// Example request body (JSON):
//
//	{
//	    "username": "johndoe",
//	    "password": "Password123!",
//	    "session_token": "existing-session-token"
//	}
//
// Note: The credentials are optional if assion token were provided and vice-versa.
// If the session token validation fails, it will use the credentials information
// if it was provided along side with the token.
//
// Example response body (JSON):
//
//	{
//	    "new_session_token": "new-session-token"
//	}
//
// Possible error responses:
//   - 401 Unauthorized: If the provided session token or credentials are invalid.
//   - 500 Internal Server Error: If there is an error reading the request body or unmarshaling JSON.
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

	if sessionToken != "" {
		newSessionToken, err := ac.Database.AuthorizeUserWithSessionToken(sessionToken)
		if err != nil {
			if err = ac.updateSessionTokenWithCredentials(w, username, password); err != nil {
				utils.WriteGenericJsonError(w, http.StatusUnauthorized,
					errors.New(INVALID_ACCESS_TOKEN_ERROR))
			}
			return
		}

		response := data.UpdateUserSessionTokenResponse{
			NewSessionToken: newSessionToken,
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

	response := data.UpdateUserSessionTokenResponse{
		NewSessionToken: newSessionToken,
	}

	responseJsonBytes, err := json.Marshal(response)
	if err != nil {
		return err
	}

	fmt.Fprint(w, string(responseJsonBytes))

	return nil
}
