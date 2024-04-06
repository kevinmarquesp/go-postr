package router

import (
	"net/http"
)

func InitRouter(port string) {
	app := http.NewServeMux()

	app.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("www/js"))))

	app.HandleFunc("/", renderIndexView)
	app.HandleFunc("/signup", renderSignupView)

	app.HandleFunc("/search/user", searchUsernameController)
	app.HandleFunc("/get/articles", getRecentArticlesController)
	app.HandleFunc("/validate/username", usernameValidationController)
	app.HandleFunc("POST /insert/user", createNewUserController)

	server := http.Server{
		Addr:    port,
		Handler: app,
	}

	server.ListenAndServe()
}

