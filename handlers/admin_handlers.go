package handlers

import (
	"github.com/dominicsieli/go-server/controllers"
	"net/http"
)

func LoginAdminHandler(response http.ResponseWriter, request *http.Request) {
	controllers.LoginAdminController(response, request)
}

func CreateAdminHandler(response http.ResponseWriter, request *http.Request) {
	controllers.CreateAdminController(response, request)
}
