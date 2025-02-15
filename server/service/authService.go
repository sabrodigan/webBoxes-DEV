package service

import (
	"fmt"
	"github.com/sabrodigan/webboxes/dto"
	"github.com/sabrodigan/webboxes/model"
	repo "github.com/sabrodigan/webboxes/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type IAuthService interface {
	Login(loginDto dto.LoginDto, sessionContext mongo.SessionContext) (map[string]interface{}, error)
}

type AuthService struct {
	repository  repo.Repository
	userService IUserService
}

func (as *AuthService) Login(loginDto dto.LoginDto, sessionContext mongo.SessionContext) (map[string]interface{}, error) {
	res, err := as.userService.FineOneUserByEmail(loginDto.Email, sessionContext)
	if err != nil {
		return nil, err
	}
	user := res.(model.User)
	if user.Password == loginDto.Password {
		return map[string]any{
			"user":  user,
			"token": "",
		}, nil
	}
	return nil, fmt.Errorf("invalid credentials")
}

func GetAuthService(repository repo.Repository, userService IUserService) IAuthService {
	return &AuthService{
		repository:  repository,
		userService: userService,
	}
}
