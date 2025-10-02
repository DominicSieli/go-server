package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/dominicsieli/go-server/services/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	address  string
	database *sql.DB
}

func NewAPIServer(address string, database *sql.DB) *APIServer {
	return &APIServer{
		address:  address,
		database: database,
	}
}

func (apiServer *APIServer) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(apiServer.database)

	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subRouter)

	log.Println("Listen on", apiServer.address)
	return http.ListenAndServe(apiServer.address, router)
}
