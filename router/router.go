package router

import (
	"log"
	"net/http"
	"time"
)

type wrappedWriter struct {
	http.ResponseWriter
	status int
}

func (ww *wrappedWriter) WriteHeader(status int) {
	ww.ResponseWriter.WriteHeader(status)

	ww.status = status
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		ww := &wrappedWriter{
			ResponseWriter: w,
			status:         http.StatusOK,
		}

		next.ServeHTTP(ww, r)

		log.Println("\t", time.Since(start), "\t", ww.status, r.Method, "\t", r.URL.Path)
	})
}

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

