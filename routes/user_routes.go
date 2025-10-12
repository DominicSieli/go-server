package routes

import (
	"github.com/dominicsieli/go-server/handlers"
	"net/http"
)

func RegisterUserRoutes(router *http.ServeMux) {
	router.HandleFunc(Post(LOGIN), handlers.LoginHandler)
	router.HandleFunc(Post(REGISTER), handlers.RegisterHandler)
}
