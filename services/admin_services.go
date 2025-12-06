package services

import (
	"github.com/dominicsieli/go-server/crud"
	"net/http"
)

func LoginAdminService(response http.ResponseWriter, request *http.Request) {
	crud.LoginAdmin(response, request)
}

func CreateAdminService(response http.ResponseWriter, request *http.Request) {
	crud.CreateAdmin(response, request)
}
