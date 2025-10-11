package main

import (
	"log"

	"github.com/dominicsieli/go-server/api"
	"github.com/dominicsieli/go-server/config"
	"github.com/dominicsieli/go-server/database"
)

func main() {
	db, err := database.CreateDatabase()

	if err != nil {
		log.Fatal(err)
	}

	database.InitializeDatabase(db)

	server := api.CreateAPIServer(config.Envs.Port, db)

	err = server.Run()

	if err != nil {
		log.Fatal(err)
	}
}
