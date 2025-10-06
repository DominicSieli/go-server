package database

import (
	"database/sql"
	"log"

	"github.com/dominicsieli/go-server/config"
	"github.com/go-sql-driver/mysql"
)

func CreateDatabase() (*sql.DB, error) {
	config := mysql.Config{
		User:				  config.Envs.DBUser,
		Passwd:				  config.Envs.DBPassword,
		Addr:				  config.Envs.DBAddress,
		DBName:				  config.Envs.DBName,
		Net:				  "tcp",
		AllowNativePasswords: true,
		ParseTime:			  true,
	}

	database, err := sql.Open("mysql", config.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	return database, nil
}

func InitializeDatabase(database *sql.DB) {
	err := database.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Database: Connection Successful!")
}
