package routes

import (
	"github.com/dominicsieli/go-server/handlers"
	"net/http"
)

func RegisterUserRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /login", handlers.UserLoginHandler)
	router.HandleFunc("POST /register", handlers.UserRegisterHandler)
}
