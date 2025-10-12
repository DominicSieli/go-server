package controllers

import (
	"github.com/dominicsieli/go-server/services"
	"net/http"
)

func LoginController(response http.ResponseWriter, request *http.Request) {
	services.LoginService(response, request)
}

func RegisterController(response http.ResponseWriter, request *http.Request) {
	services.RegisterService(response, request)
}
