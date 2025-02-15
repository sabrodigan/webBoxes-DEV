package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sabrodigan/webboxes/controller"
	repo "github.com/sabrodigan/webboxes/repository"
	"github.com/sabrodigan/webboxes/service"
	"github.com/sabrodigan/webboxes/utils"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	repository := repo.GetRepository().UserRepository
	userService := service.GetUserService(repository)
	responseService := utils.GetResponseService()
	userController := controller.GetUserController(userService, *responseService)

	router.POST(
		"/add",
		userController.CreateUser,
	)

}
