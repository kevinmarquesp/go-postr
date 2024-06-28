package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kevinmarquesp/go-postr/internal/data"
)

func WriteJsonError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)

	jsonError := &data.ErrorResponse{
		Status:     status,
		StatusText: http.StatusText(status),
		Error:      err.Error(),
	}

	errorBytes, _ := json.Marshal(jsonError)
	errorString := string(errorBytes)

	fmt.Fprint(w, errorString)
}
