package handlers

import (
	"github.com/dominicsieli/go-server/controllers"
	"net/http"
)

func LoginHandler(response http.ResponseWriter, request *http.Request) {
	controllers.LoginController(response, request)
}

func RegisterHandler(response http.ResponseWriter, request *http.Request) {
	controllers.RegisterController(response, request)
}
