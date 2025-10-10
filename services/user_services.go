package services

import (
	"github.com/dominicsieli/go-server/data"
	"net/http"
)

func UserLoginService(response http.ResponseWriter, request *http.Request) {
	data.Create(response, request)
}

func UserRegisterService(response http.ResponseWriter, request *http.Request) {
	data.Create(response, request)
}
