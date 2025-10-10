package controllers

import (
	"github.com/dominicsieli/go-server/services"
	"net/http"
)

func UserLoginController(response http.ResponseWriter, request *http.Request) {
	services.UserLoginService(response, request)
}

func UserRegisterController(response http.ResponseWriter, request *http.Request) {
	services.UserRegisterService(response, request)
}
