package router

import (
	"net/http"

	"github.com/charmbracelet/log"
)

func InitRouter(port string) {
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("www/js"))))

	http.HandleFunc("/", renderIndexView)
	http.HandleFunc("/signup", renderSignupView)

	http.HandleFunc("/search/user", searchUsernameController)
	http.HandleFunc("/get/articles", getRecentArticlesController)
	http.HandleFunc("/validate/username", usernameValidationController)
	http.HandleFunc("/insert/user", createNewUserController)

	log.Info("Listening to", "url", "http://localhost" + port)
	http.ListenAndServe(port, nil)
}

