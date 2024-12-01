package responses

import "github.com/gin-gonic/gin"

type CommonResponse struct {
	Message string
	Data    interface{}
}

func RespondWithError(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, CommonResponse{
		Message: "Error",
		Data:    data,
	})
}
