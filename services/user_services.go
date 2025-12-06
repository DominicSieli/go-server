package services

import (
	"github.com/dominicsieli/go-server/crud"
	"net/http"
)

func LoginUserService(response http.ResponseWriter, request *http.Request) {
	crud.LoginUser(response, request)
}

func CreateUserService(response http.ResponseWriter, request *http.Request) {
	crud.CreateUser(response, request)
}
