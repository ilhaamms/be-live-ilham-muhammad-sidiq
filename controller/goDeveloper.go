package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilhaamms/backend-live/models/response"
	"github.com/ilhaamms/backend-live/service"
)

type GoDeveloperController interface {
	GoDeveloper(ctx *gin.Context)
}

type GoDeveloperControllers struct {
	goDeveloperService service.GoDeveloperService
}

func NewGoDeveloperController(goDeveloperService service.GoDeveloperService) GoDeveloperController {
	return &GoDeveloperControllers{
		goDeveloperService: goDeveloperService,
	}
}

func (c *GoDeveloperControllers) GoDeveloper(ctx *gin.Context) {
	result, err := c.goDeveloperService.GoDeveloper()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Success{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       result,
	})
}
