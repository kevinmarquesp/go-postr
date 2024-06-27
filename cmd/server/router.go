package main

import (
	"errors"
	"fmt"
	"net/http"
)

func initServerRouter(port string) error {
	if port == "" {
		return errors.New("The port environment was not specified.")
	}

	apiRouter := http.NewServeMux()

	apiRouter.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		fmt.Fprint(w, `{ "message": "Hello world!" }`)
	})

	router := http.NewServeMux()

	router.Handle("/api/", http.StripPrefix("/api", apiRouter))

	server := http.Server{
		Addr:    ":" + port,
		Handler: MiddlewareHandler(router),
	}

	return server.ListenAndServe()
}
