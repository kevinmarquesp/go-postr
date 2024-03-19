package router

import (
	"net/http"

	"github.com/charmbracelet/log"
)

func InitRouter(port string) {
	http.HandleFunc("/", RenderIndexController)
	http.HandleFunc("/search/user", SearchUsernameController)

	log.Info("Listening to", "url", "http://localhost" + port)
	http.ListenAndServe(port, nil)
}

