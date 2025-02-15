package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sabrodigan/webboxes/model"
	"github.com/sabrodigan/webboxes/service"
	"github.com/sabrodigan/webboxes/utils"
)

type UserController struct {
	userService     service.IUserService
	responseService utils.ResponseService
}

func (uc *UserController) CreateUser(ctx *gin.Context) {

	var dto model.UserCreateDto

	if err := ctx.ShouldBindJSON(&dto); err != nil {

		ctx.Error(fmt.Errorf("400::%s::%s::%v", "Invalid request", err.Error(), err))
		return
	}
	data, err := uc.userService.CreateUser(dto, nil)

	if err != nil {
		ctx.Error(err)
		return
	}

	uc.responseService.Success(ctx, 201, data, "successfully saved")

}
func GetUserController(userService service.IUserService, responseService utils.ResponseService) *UserController {
	return &UserController{
		userService:     userService,
		responseService: responseService,
	}
}
