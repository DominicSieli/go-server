package controllers

import (
	"github.com/dominicsieli/go-server/services"
	"net/http"
)

func AdminLoginController(response http.ResponseWriter, request *http.Request) {
	services.AdminLoginService(response, request)
}

func AdminRegisterController(response http.ResponseWriter, request *http.Request) {
	services.AdminRegisterService(response, request)
}
