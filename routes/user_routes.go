package routes

import (
	"github.com/dominicsieli/go-server/handlers"
	"net/http"
)

const (
	USER_LOGIN    = "user/login"
	USER_REGISTER = "user/register"
)

func RegisterUserRoutes(router *http.ServeMux) {
	router.HandleFunc(Post(USER_LOGIN), handlers.UserLoginHandler)
	router.HandleFunc(Post(USER_REGISTER), handlers.UserRegisterHandler)
}
