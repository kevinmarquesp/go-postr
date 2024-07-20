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
	// apiRouter.HandleFunc("GET /user/{username}", apiUserHandler.FetchUserDataByUsername)
	// apiRouter.HandleFunc("PUT /user", apiUserHandler.UpdateProfileDetails)
	// apiRouter.HandleFunc("DELETE /user", apiUserHandler.RemoveRegisteredUser)
	// apiRouter.HandleFunc("GET /users", apiUserHandler.FetchUsersList)

	router := http.NewServeMux()

	router.HandleFunc("/tailwind.min.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		http.ServeFile(w, r, "./dist/css/tailwind.min.css")
	})

	router.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./www/css"))))

	pagesHandler := controllers.NewPagesHanlder()

	router.HandleFunc("/", pagesHandler.RenderHomePage)
	// router.HandleFunc("/u/{username}", pagesHandler.RenderProfilePage)
	// router.HandleFunc("/publication/{id}", pagesHandler.RenderProfilePage)

	// authHandler := constrollers.NewAuthHanlder()

	// router.HandleFunc("/register", authHandler.RegisterNewUser)
	// router.HandleFunc("/quit", authHandler.RegisterNewUser)
	// router.HandleFunc("/delete", authHandler.RegisterNewUser)

	router.Handle("/api/v1/", http.StripPrefix("/api/v1", apiRouter))

	// Server initialization.

	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	return server.ListenAndServe()
}
