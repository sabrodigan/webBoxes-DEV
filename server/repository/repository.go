package repository

import (
	"log"

	"github.com/sabrodigan/webboxes/config"
)

type Repository struct {
	UserRepository IMongoRepository
}

// GetRepository initializes the Repository
func GetRepository() *Repository {
	// Retrieve the database name and handle any potential error
	dbName, err := config.GetEnvProperty("dataBaseName")

	if err != nil {
		// Handle error appropriately, for example, logs and exits the application
		log.Fatalf("Error fetching database name: %v", err)
	}
	users, err := config.GetEnvProperty("users")
	if err != nil {
		log.Fatalf("Error fetching users collection: %v", err)
	}
	// Pass the dbName and collection to the GetMongoRepository function
	return &Repository{
		UserRepository: GetMongoRepository(dbName, users),
	}
}
