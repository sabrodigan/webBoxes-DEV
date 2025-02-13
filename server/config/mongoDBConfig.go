package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var dbClient *mongo.Client
var DEFAULTDBNAME = "storage"

func init() {
	LoadEnvironmentVariable()
}

func InitDatabase() (*mongo.Client, error) {
	dbURL, _ := GetEnvProperty("databaseURL")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURL))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	return client, nil

}
func GetDatabaseCollection(dbName *string, collectionName string) (*mongo.Collection, error) {
	if *dbName == "" {
		*dbName = DEFAULTDBNAME
	}

	collection := dbClient.Database(*dbName).Collection(collectionName)
	return collection, nil
}
func InitializeDB() (*mongo.Client, error) {
	db, err := InitDatabase()
	if err != nil {
		log.Fatalf("failed to connect to the db %v", err)
	}
	return db, nil
}
