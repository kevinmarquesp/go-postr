package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/kevinmarquesp/go-postr/internal/data"
	"github.com/kevinmarquesp/go-postr/internal/models"
	"github.com/kevinmarquesp/go-postr/internal/utils"
)

type UsersController struct {
	Database models.GenericDatabaseProvider
}

func (us UsersController) UpdateUserProfileDetails(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	rawBody, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteGenericJsonError(w, http.StatusInternalServerError, err)
		return
	}

	var body data.UpdateUserProfileDetailsBody

	if err = json.Unmarshal(rawBody, &body); err != nil {
		utils.WriteGenericJsonError(w, http.StatusInternalServerError, err)
		return
	}

	userPublicId := r.PathValue("userPublicId")
	sessionToken := strings.Trim(body.SessionToken, " ")

	log.Infof("Public Id:\t\t%s", userPublicId)
	log.Infof("Session Token:\t%s", sessionToken)

	// TODO: Create a function that recieves the body fields to update the user.
	// TODO: Get the updated user profile fields from that function.
	// TODO: Format a JSON response, then send it to the final user.
}
