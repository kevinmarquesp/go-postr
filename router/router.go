package router

import (
	"net/http"

	"github.com/charmbracelet/log"
)

func InitRouter(port string) {
	http.HandleFunc("/", renderIndexController)
	http.HandleFunc("/search/user", searchUsernameController)
	http.HandleFunc("/get/articles", getRecentArticlesController)

	log.Info("Listening to", "url", "http://localhost" + port)
	http.ListenAndServe(port, nil)
}

