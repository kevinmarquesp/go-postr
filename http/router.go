package http

import (
	"net/http"
)

func SetupRoutes() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("www/css"))))

	http.HandleFunc("/", homePageController)
	http.HandleFunc("/signup", signupPageController)
}

func InitializeRouter(port string) error {
	err := http.ListenAndServe(port, nil)
	if err != nil {
		return err
	}

	return nil
}
