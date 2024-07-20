package models

import (
	"encoding/json"
	"net/http"
)

type JsonError struct {
	Status     int    `json:"status"`
	StatusText string `json:"statusText"`
	Message    string `json:"message"`
	Location   string `json:"location"`
	Error      string `json:"error"`
}

func WriteHttpJsonError(w http.ResponseWriter, status int, err error, location, message string) {
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(JsonError{
		Status:     status,
		StatusText: http.StatusText(status),
		Message:    message,
		Location:   location,
		Error:      err.Error(),
	})
}
