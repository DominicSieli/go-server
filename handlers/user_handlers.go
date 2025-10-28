package handlers

import (
	"github.com/dominicsieli/go-server/controllers"
	"net/http"
)

func UserLoginHandler(response http.ResponseWriter, request *http.Request) {
	controllers.UserLoginController(response, request)
}

func UserRegisterHandler(response http.ResponseWriter, request *http.Request) {
	controllers.UserRegisterController(response, request)
}
