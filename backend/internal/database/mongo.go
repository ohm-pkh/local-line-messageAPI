package database

import (
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func New() *mongo.Database {
	uri := os.Getenv("MONGODB_URI")
	databaseName := os.Getenv("DATABASE_NAME")
	client, err := mongo.Connect(options.Client().
		ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	db := client.Database(databaseName)
	return db
}
