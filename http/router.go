package http

import (
	"fmt"
	"net/http"
)

func InitializeRouter(port string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hello world!</h1>")
	})

	err := http.ListenAndServe(port, nil)
	if err != nil {
		return err
	}

	return nil
}
