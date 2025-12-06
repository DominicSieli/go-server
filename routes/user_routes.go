package routes

import (
	"github.com/dominicsieli/go-server/handlers"
	"net/http"
)

const (
	USER_LOGIN  = "user/login"
	USER_CREATE = "user/create"
)

func RegisterUserRoutes(router *http.ServeMux) {
	router.HandleFunc(Post(USER_LOGIN), handlers.LoginUserHandler)
	router.HandleFunc(Post(USER_CREATE), handlers.CreateUserHandler)
}
