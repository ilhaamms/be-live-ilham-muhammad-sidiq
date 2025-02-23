package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilhaamms/backend-live/helper"
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
		helper.ResponseJsonError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	helper.ResponseJsonSuccess(ctx, http.StatusOK, "Success", result)
}
