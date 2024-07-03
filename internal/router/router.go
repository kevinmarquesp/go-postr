package router

import (
	"errors"
	"net/http"

	"github.com/kevinmarquesp/go-postr/internal/api"
	"github.com/kevinmarquesp/go-postr/internal/models"
)

func InitRouter(port string, db models.GenericDatabaseProvider) error {
	if port == "" {
		return errors.New("The port environment was not specified.")
	}

	// API router.

	apiRouter := http.NewServeMux()

	authController := api.AuthController{Database: db}
	usersController := api.UsersController{Database: db}
	//publicController := api.PublicController{Database: db}

	apiRouter.HandleFunc("POST /auth/register", authController.RegisterNewUser)
	apiRouter.HandleFunc("PUT /auth/refresh", authController.RefreshUserSessionToken)
	//apiRouter.HandleFunc("DELETE /auth/quit", authController.RemoveUser)
	apiRouter.HandleFunc("PUT /settings/profile/{userPublicId}", usersController.UpdateUserProfileDetails)
	//apiRouter.HandleFunc("GET /profile/{username}", publicController.GetPublicUserInformation)
	//apiRouter.HandleFunc("POST /follow/{userPublicId}", usersController.FollowUserByPublicId)
	//apiRouter.HandleFunc("GET /users", publicController.GetUsersSummaryList)

	// Global router.

	router := http.NewServeMux()

	router.Handle("/api/", http.StripPrefix("/api", apiRouter))

	// Application server setup.

	server := http.Server{
		Addr:    ":" + port,
		Handler: MiddlewareHandler(router),
	}

	return server.ListenAndServe()
}
