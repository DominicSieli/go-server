package services

import (
	"github.com/dominicsieli/go-server/data"
	"net/http"
)

func LoginService(response http.ResponseWriter, request *http.Request) {
	data.Login(response, request)
}

func RegisterService(response http.ResponseWriter, request *http.Request) {
	data.Register(response, request)
}
