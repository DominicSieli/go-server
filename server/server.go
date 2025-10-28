package server

import (
	"log"
	"net/http"

	"github.com/dominicsieli/go-server/config"
	"github.com/dominicsieli/go-server/database"
	"github.com/dominicsieli/go-server/middleware"
	"github.com/dominicsieli/go-server/routes"
)

func StartServer() {
	router := http.NewServeMux()
	routes.RegisterUserRoutes(router)

	db, err := database.CreateDatabase()

	if err != nil {
		log.Fatal(err)
	}

	database.InitializeDatabase(db)

	stack := middleware.CreateStack(
		middleware.Logging,
	)

	server := http.Server{
		Addr:    config.Envs.Port,
		Handler: stack(router),
	}

	log.Println("Listening on", server.Addr)

	server.ListenAndServe()
}
