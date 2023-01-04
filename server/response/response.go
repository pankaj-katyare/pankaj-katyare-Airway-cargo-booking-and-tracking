package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Response(context *gin.Context, statusCode int, data interface{}) {
	context.JSON(statusCode, data)
}

func ResponseBytes(context *gin.Context, statusCode int, data []byte) {
	context.Data(statusCode, "application/json", data)
}

func SuccessResponse(context *gin.Context, data interface{}) {
	Response(context, http.StatusOK, data)
}

func ErrorResponse(context *gin.Context, statusCode int, message string) {
	Response(context, statusCode, Error{Code: statusCode, Message: message})
}
