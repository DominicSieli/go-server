package routes

import (
	"github.com/dominicsieli/go-server/handlers"
	"net/http"
)

func RegisterUserRoutes(router *http.ServeMux) {
	router.HandleFunc(Post(USER_LOGIN), handlers.UserLoginHandler)
	router.HandleFunc(Post(USER_REGISTER), handlers.UserRegisterHandler)
}
