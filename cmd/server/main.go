package main

import (
	"fmt"
	"net/http"
)

func main() {
	app := http.NewServeMux()

	app.HandleFunc("GET /", func(w http.ResponseWriter, _r *http.Request) {
		fmt.Fprint(w, "<h1>Hello world! From Go 1.22.1</h1>")
	})

	server := http.Server{
		Addr: ":8080",
		Handler: app,
	}

	server.ListenAndServe()
}
