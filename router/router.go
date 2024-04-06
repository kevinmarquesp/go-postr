package router

import (
	"net/http"

	"github.com/charmbracelet/log"
)

func InitRouter(port string) {
	http.Handle("GET /js/", http.StripPrefix("/js/", http.FileServer(http.Dir("www/js"))))

	http.HandleFunc("GET /home", renderIndexView)
	http.HandleFunc("GET /signup", renderSignupView)

	http.HandleFunc("GET /search/user", searchUsernameController)
	http.HandleFunc("GET /get/articles", getRecentArticlesController)
	http.HandleFunc("GET /validate/username", usernameValidationController)
	http.HandleFunc("POST /insert/user", createNewUserController)

	log.Info("Listening to", "url", "http://localhost" + port)
	http.ListenAndServe(port, nil)
}

