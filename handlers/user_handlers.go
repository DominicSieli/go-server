package handlers

import (
	"github.com/dominicsieli/go-server/controllers"
	"net/http"
)

func LoginUserHandler(response http.ResponseWriter, request *http.Request) {
	controllers.LoginUserController(response, request)
}

func CreateUserHandler(response http.ResponseWriter, request *http.Request) {
	controllers.CreateUserController(response, request)
}
