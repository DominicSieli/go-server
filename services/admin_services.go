package services

import (
	"github.com/dominicsieli/go-server/data"
	"net/http"
)

func AdminLoginService(response http.ResponseWriter, request *http.Request) {
	data.Login(response, request)
}

func AdminRegisterService(response http.ResponseWriter, request *http.Request) {
	data.Register(response, request)
}
