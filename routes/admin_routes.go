package routes

import (
	"github.com/dominicsieli/go-server/handlers"
	"net/http"
)

const (
	ADMIN_LOGIN    = "admin/login"
	ADMIN_REGISTER = "admin/register"
)

func RegisterAdminRoutes(router *http.ServeMux) {
	router.HandleFunc(Post(ADMIN_LOGIN), handlers.AdminLoginHandler)
	router.HandleFunc(Post(ADMIN_REGISTER), handlers.AdminRegisterHandler)
}
