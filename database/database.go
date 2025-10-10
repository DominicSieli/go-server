package database

import (
	"context"
	"log"

	"github.com/dominicsieli/go-server/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateDatabase() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(config.Envs.DBURI)
	database, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	return database, nil
}

func InitializeDatabase(database *mongo.Client) {
	err := database.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Database: Connection Successful!")
}
