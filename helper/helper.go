package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ilhaamms/backend-live/models/response"
)

var Validate = validator.New()

func ResponseJsonSuccess(ctx *gin.Context, statusCode int, message string, data any) {
	ctx.JSON(http.StatusOK, response.Success{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	})
}

func ResponseJsonError(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(statusCode, response.Error{
		StatusCode: statusCode,
		Message:    message,
	})
}
