package routes

import (
	"github.com/dominicsieli/go-server/handlers"
	"net/http"
)

const (
	ADMIN_LOGIN  = "admin/login"
	ADMIN_CREATE = "admin/create"
)

func RegisterAdminRoutes(router *http.ServeMux) {
	router.HandleFunc(Post(ADMIN_LOGIN), handlers.LoginAdminHandler)
	router.HandleFunc(Post(ADMIN_CREATE), handlers.CreateAdminHandler)
}
