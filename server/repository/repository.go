package repository

import "github.com/sabrodigan/webboxes/config"

type Repository struct {
	UserRepository IMongoRepository
}

func GetRepository() *Repository {
	return &Repository{
		UserRepository: GetMongoRepository(config.GetEnvProperty("databaseName"), "user"),
	}
}
