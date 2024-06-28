package utils

import (
	"fmt"
	"net/http"
)

func WriteJsonError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	fmt.Fprintf(w, `{ "error": "%s" }`, err)
}
