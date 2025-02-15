package service

import (
	"github.com/sabrodigan/webboxes/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserService interface {
	CreateUser(data interface{}, sessionContext mongo.SessionContext) (interface{}, error)
}

type UserService struct {
	repository repository.IMongoRepository
}

func (us *UserService) CreateUser(data interface{}, sessionContext mongo.SessionContext) (interface{}, error) {
	return us.repository.Create(data, sessionContext)
}

func GetUserService(repository repository.IMongoRepository) IUserService {
	return &UserService{
		repository: repository,
	}
}
