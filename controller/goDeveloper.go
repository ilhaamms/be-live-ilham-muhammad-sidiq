package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func (controller *GoDeveloperControllers) GoDeveloper(ctx *gin.Context) {
	result, err := controller.goDeveloperService.GoDeveloper()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
