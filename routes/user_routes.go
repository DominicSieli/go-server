package routes

import (
	"github.com/dominicsieli/go-server/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/login", handlers.UserLoginHandler).Methods(http.MethodPost)
	router.HandleFunc("/register", handlers.UserRegisterHandler).Methods(http.MethodPost)
}
