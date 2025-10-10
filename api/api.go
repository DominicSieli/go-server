package api

import (
	"log"
	"net/http"

	"github.com/dominicsieli/go-server/routes"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type APIServer struct {
	address  string
	database *mongo.Client
}

func CreateAPIServer(address string, database *mongo.Client) *APIServer {
	return &APIServer{
		address:  address,
		database: database,
	}
}

func (apiServer *APIServer) Run() error {
	router := mux.NewRouter()

	routes.RegisterUserRoutes(router)

	log.Println("Listening on", apiServer.address)

	return http.ListenAndServe(apiServer.address, router)
}
