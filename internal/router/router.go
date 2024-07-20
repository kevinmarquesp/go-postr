package router

import (
	"errors"
	"net/http"

	"github.com/kevinmarquesp/go-postr/internal/controllers"
	"github.com/kevinmarquesp/go-postr/internal/repositories"
	"github.com/kevinmarquesp/go-postr/internal/services"
)

func StartRouter(port string) error {
	if port == "" {
		return errors.New("Empty port value error.")
	}

	// Enabled repositories.

	userRepository, err := repositories.NewSqliteUserRepository("./tmp/database.sqlite3")
	if err != nil {
		return err
	}

	// Enalbed services.

	authenticationService := services.NewGopostrAuthenticationService(userRepository)

	// Routes definition.

	apiRouter := http.NewServeMux()

	apiUserHandler := controllers.NewApiUserHandler(authenticationService)

	apiRouter.HandleFunc("POST /user", apiUserHandler.RegisterNewUserWithCredentials)
	apiRouter.HandleFunc("GET /user/{username}", apiUserHandler.FetchUserDataByUsername)
	apiRouter.HandleFunc("PUT /user", apiUserHandler.UpdateProfileDetails)
	apiRouter.HandleFunc("DELETE /user", apiUserHandler.RemoveRegisteredUser)

	router := http.NewServeMux()

	router.Handle("/api/v1/", http.StripPrefix("/api/v1", apiRouter))

	// Server initialization.

	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	return server.ListenAndServe()
}
