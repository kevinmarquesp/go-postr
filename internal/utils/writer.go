package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kevinmarquesp/go-postr/internal/data"
)

// This function sets the HTTP status code, marshals the error message into a JSON format,
// and writes it to the response writer. It is useful for sending structured error responses
// on the API routes.
//
// Example of a JSON response:
//
//	{
//	    "StatusText": "Internal Server Error",
//	    "Error": "something went wrong"
//	}
func WriteGenericJsonError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)

	jsonError := &data.GenericErrorResponse{
		StatusText: http.StatusText(status),
		Error:      err.Error(),
	}

	errorBytes, _ := json.Marshal(jsonError)
	errorString := string(errorBytes)

	fmt.Fprint(w, errorString)
}
