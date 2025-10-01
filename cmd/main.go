package main

import (
	"database/sql"
	"log"

	"github.com/dominicsieli/go-server-template/cmd/api"
	"github.com/dominicsieli/go-server-template/config"
	"github.com/dominicsieli/go-server-template/database"
	"github.com/go-sql-driver/mysql"
)

func main() {
	database, err := database.NewMySQLStorage(mysql.Config{
		User: config.Envs.DBUser,
		Passwd: config.Envs.DBPassword,
		Addr: config.Envs.DBAddress,
		DBName: config.Envs.DBName,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	})

	if err != nil {
		log.Fatal(err)
	}

	//initStorage(database)

	server := api.NewAPIServer(":8080", database)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(database *sql.DB) {
	err := database.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Database: Connection Successful!")
}