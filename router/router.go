package router

import (
	"net/http"

	"github.com/charmbracelet/log"
)

func InitRouter(port string) {
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("www/js"))))

	http.HandleFunc("/", renderIndexController)
	http.HandleFunc("/signup", renderSignupController)
	http.HandleFunc("/search/user", searchUsernameController)
	http.HandleFunc("/get/articles", getRecentArticlesController)
	http.HandleFunc("/validate/username", usernameValidationController)

	log.Info("Listening to", "url", "http://localhost" + port)
	http.ListenAndServe(port, nil)
}

