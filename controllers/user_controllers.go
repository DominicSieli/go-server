package controllers

import (
	"github.com/dominicsieli/go-server/services"
	"net/http"
)

func LoginUserController(response http.ResponseWriter, request *http.Request) {
	services.LoginUserService(response, request)
}

func CreateUserController(response http.ResponseWriter, request *http.Request) {
	services.CreateUserService(response, request)
}
