package services

import (
	"github.com/dominicsieli/go-server/crud"
	"net/http"
)

func UserLoginService(response http.ResponseWriter, request *http.Request) {
	crud.LoginUser(response, request)
}

func UserRegisterService(response http.ResponseWriter, request *http.Request) {
	crud.RegisterUser(response, request)
}
