package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	router := r.Group("/")

	userRoutes := router.Group("/users")
	{
		RegisterUserRoutes(userRoutes)

	}
	authRoutes := router.Group("/auth")
	{
		RegisterAuthRoutes(authRoutes)
	}
}
