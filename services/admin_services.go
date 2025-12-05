package services

import (
	"github.com/dominicsieli/go-server/crud"
	"net/http"
)

func AdminLoginService(response http.ResponseWriter, request *http.Request) {
	crud.AdminLogin(response, request)
}

func AdminRegisterService(response http.ResponseWriter, request *http.Request) {
	crud.AdminRegister(response, request)
}
