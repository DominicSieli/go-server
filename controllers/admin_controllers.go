package controllers

import (
	"github.com/dominicsieli/go-server/services"
	"net/http"
)

func LoginAdminController(response http.ResponseWriter, request *http.Request) {
	services.LoginAdminService(response, request)
}

func CreateAdminController(response http.ResponseWriter, request *http.Request) {
	services.CreateAdminService(response, request)
}
