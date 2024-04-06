package router

import (
	"net/http"
)

func InitRouter(port string) {
	app := http.NewServeMux()

	app.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("www/js"))))

	app.HandleFunc("/", renderIndexView)
	app.HandleFunc("/signup", renderSignupView)

	app.HandleFunc("/api/search/user", searchUsernameController)
	app.HandleFunc("/api/get/articles", getRecentArticlesController)
	app.HandleFunc("/api/validate/username", usernameValidationController)
	app.HandleFunc("POST /api/insert/user", createNewUserController)

	server := http.Server{
		Addr:    port,
		Handler: app,
	}

	server.ListenAndServe()
}

