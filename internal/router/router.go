package router

import (
	"qd_authentication_api/internal/handlers"
	"qd_authentication_api/internal/service"

	"github.com/gorilla/mux"
)

func SetupRoutes(authService *service.AuthService) *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/register", handlers.RegisterHandler(authService)).Methods("POST")

	return router
}
