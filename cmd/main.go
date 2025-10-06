package main

import (
	"log"

	"github.com/dominicsieli/go-server/api"
	"github.com/dominicsieli/go-server/database"
)

func main() {
	db, err := database.CreateDatabase()

	if err != nil {
		log.Fatal(err)
	}

	//database.InitializeDatabase(db)
	server := api.CreateAPIServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
