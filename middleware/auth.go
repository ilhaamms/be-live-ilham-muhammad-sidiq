package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilhaamms/backend-live/models/response"
)

func MethodNotAllowedMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, response.Error{
			StatusCode: http.StatusMethodNotAllowed,
			Message:    "Method Not Allowed",
		})
		c.Abort()
	}
}
