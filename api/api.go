package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/dominicsieli/go-server/controllers"
	"github.com/dominicsieli/go-server/repos"
	"github.com/gorilla/mux"
)

type APIServer struct {
	address  string
	database *sql.DB
}

func CreateAPIServer(address string, database *sql.DB) *APIServer {
	return &APIServer{
		address:  address,
		database: database,
	}
}

func (apiServer *APIServer) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	userRepo := repos.CreateRepo(apiServer.database)

	userHandler := controllers.CreateController(userRepo)
	userHandler.RegisterRoutes(subRouter)

	log.Println("Listening on", apiServer.address)

	return http.ListenAndServe(apiServer.address, router)
}
