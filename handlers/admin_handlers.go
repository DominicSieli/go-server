package handlers

import (
	"github.com/dominicsieli/go-server/controllers"
	"net/http"
)

func AdminLoginHandler(response http.ResponseWriter, request *http.Request) {
	controllers.AdminLoginController(response, request)
}

func AdminRegisterHandler(response http.ResponseWriter, request *http.Request) {
	controllers.AdminRegisterController(response, request)
}
