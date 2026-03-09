package response

import "github.com/gin-gonic/gin"

type APIResponse struct {
	Status       bool        `json:"status"`
	StatusCode   int         `json:"status_code"`
	Data         interface{} `json:"data,omitempty"`
	Message      string      `json:"message,omitempty"`
	ErrorMessage string      `json:"error_message,omitempty"`
}

func Success(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, APIResponse{
		Status:     true,
		StatusCode: statusCode,
		Data:       data,
		Message:    message,
	})
}

func Error(c *gin.Context, statusCode int, errorMessage string) {
	c.JSON(statusCode, APIResponse{
		Status:       false,
		StatusCode:   statusCode,
		ErrorMessage: errorMessage,
	})
}

func AbortError(c *gin.Context, statusCode int, errorMessage string) {
	c.AbortWithStatusJSON(statusCode, APIResponse{
		Status:       false,
		StatusCode:   statusCode,
		ErrorMessage: errorMessage,
	})
}
