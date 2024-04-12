package router

import "net/http"

func InitRouter(port string) {
	app := http.NewServeMux()

	app.Handle("GET /js/", http.StripPrefix("/js/", http.FileServer(http.Dir("www/js"))))

	app.HandleFunc("GET /", renderIndexView)
	app.HandleFunc("GET /signup", renderSignupView)

	app.HandleFunc("GET /api/search/user", searchUsernameController)
	app.HandleFunc("GET /api/get/articles", getRecentArticlesController)
	app.HandleFunc("GET /api/validate/u/{username}", usernameValidationController)
	app.HandleFunc("POST /api/insert/user", createNewUserController)

	server := http.Server{
		Addr:    port,
		Handler: loggingMiddleware(app),
	}

	server.ListenAndServe()
}

