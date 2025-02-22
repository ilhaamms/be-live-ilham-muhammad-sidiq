package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilhaamms/backend-live/service"
)

type LanguageController interface {
	GetLanguage(ctx *gin.Context)
}

type LanguageControllers struct {
	languageService service.LanguageService
}

func NewLanguageController(languageService service.LanguageService) LanguageController {
	return &LanguageControllers{
		languageService: languageService,
	}
}

func (controller *LanguageControllers) GetLanguage(ctx *gin.Context) {
	result, err := controller.languageService.FecthLanguage()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
