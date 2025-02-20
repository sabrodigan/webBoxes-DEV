package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sabrodigan/webboxes/controller"
	repo "github.com/sabrodigan/webboxes/repository"
	"github.com/sabrodigan/webboxes/service"
	"github.com/sabrodigan/webboxes/utils"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	repository := repo.GetRepository()
	userService := service.GetUserService(*repository)
	authService := service.GetAuthService(*repository, userService)
	responseService := utils.GetResponseService()
	authController := controller.GetAuthController(authService, *responseService)

	router.POST(
		"/login",
		authController.Login,
	)
}
