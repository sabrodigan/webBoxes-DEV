package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sabrodigan/webboxes/config"
	"github.com/sabrodigan/webboxes/routes"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	fmt.Println("Starting application...")

	app := gin.Default()
	app.Use(gin.Recovery())

	app.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message":    "pong",
			"statusCode": 200,
		})
	})
	routes.RegisterRoutes(app)
	
	fmt.Println("Our application is now running!")

	port, err := config.GetEnvProperty("port")
	if err != nil {
		fmt.Println("Error getting port:", err)
		return
	}

	fmt.Println("Running on port:", port)
	err = app.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		fmt.Println("Error running application:", err)
		return
	}
}
