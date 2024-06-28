package utils

import (
	"fmt"
	"net/http"
)

func WriteJsonError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)

	fmt.Fprintf(w, `{ "status": %d, "statusText": "%s", "error": "%s" }`,
		status, http.StatusText(status), err)
}
