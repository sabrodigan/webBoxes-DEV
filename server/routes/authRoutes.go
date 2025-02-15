package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sabrodigan/webboxes/controller"
	repo "github.com/sabrodigan/webboxes/repository"
	"github.com/sabrodigan/webboxes/service"
	"github.com/sabrodigan/webboxes/utils"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {
	repository := repo.GetRepository()
	userService := service.GetUserService(*repository)              // Ensure correct arguments
	authService := service.GetAuthService(*repository, userService) // Ensure correct arguments
	responseService := utils.GetResponseService()
	authController := controller.GetAuthController(authService, *responseService) // Ensure correct type

	router.POST(
		"/login",
		authController.Login,
	)
}
