package utils

import "github.com/gin-gonic/gin"

type ResponseService struct {
	data interface{}
}

func (rs *ResponseService) Success(ctx *gin.Context, statusCode int, data interface{}, message string) {
	var finalResponse = map[string]any{
		"data":       data,
		"statusCode": statusCode,
		"message":    message,
	}

	if message == "" {
		finalResponse["message"] = "success"
	}

	ctx.JSON(statusCode, finalResponse)
}

func GetResponseService() *ResponseService {
	return &ResponseService{}
}
