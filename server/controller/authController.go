package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sabrodigan/webboxes/dto"
	"github.com/sabrodigan/webboxes/service"
	"github.com/sabrodigan/webboxes/utils"
)

type AuthController struct {
	authService     service.IAuthService
	responseService utils.ResponseService
}

func (ac *AuthController) Login(ctx *gin.Context) {

	var loginDto dto.LoginDto
	if err := ctx.ShouldBindJSON(&loginDto); err != nil {
		ctx.Error(fmt.Errorf("400::%s::%s::%v", "Invalid request", err.Error(), err))
		return
	}
	data, err := ac.authService.Login(loginDto, nil)
	if err != nil {
		ctx.Error(err)
		return
	}

	ac.responseService.Success(ctx, 201, data, "successfully saved")
}

func GetAuthController(authService service.IAuthService, responseService utils.ResponseService) *AuthController {
	return &AuthController{
		authService:     authService,
		responseService: responseService,
	}
}
