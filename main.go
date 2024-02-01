package main

import (
	"go-postr/http"
	"log"
)

func main() {
	err := http.InitializeRouter(":8080")
	if err != nil {
		log.Panicln("http router initialization error:", err)
	}
}
