package http

import (
	"net/http"
)

func SetupRoutes() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("www/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("www/js"))))

	http.HandleFunc("/component/FieldValidationStatus", getFieldValidationStatusComponent)

	http.HandleFunc("/", homePageController)
	http.HandleFunc("/auth/signup", signupPageController)
	http.HandleFunc("/auth/create", createNewUserController)
	http.HandleFunc("/auth/validate/username", usernameValidationController)
}

func InitializeRouter(port string) error {
	err := http.ListenAndServe(port, nil)
	if err != nil {
		return err
	}

	return nil
}
