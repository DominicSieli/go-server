package services

import (
	"github.com/dominicsieli/go-server/crud"
	"net/http"
)

func AdminLoginService(response http.ResponseWriter, request *http.Request) {
	crud.LoginAdmin(response, request)
}

func AdminRegisterService(response http.ResponseWriter, request *http.Request) {
	crud.RegisterAdmin(response, request)
}
