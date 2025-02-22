package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ilhaamms/backend-live/helper"
	"github.com/ilhaamms/backend-live/models/entity"
	"github.com/ilhaamms/backend-live/models/response"
	"github.com/ilhaamms/backend-live/service"
)

type LanguageController interface {
	GetLanguage(ctx *gin.Context)
	AddLanguage(ctx *gin.Context)
	GetAllLanguage(ctx *gin.Context)
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

	idParam := ctx.Param("id")
	if idParam == "" {
		ctx.JSON(400, response.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "parameter id is required, please input id",
		})
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, response.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "parameter id must be a number",
		})
		return
	}

	result, err := controller.languageService.FecthLanguageByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error{
			StatusCode: http.StatusBadRequest,
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

func (controller *LanguageControllers) AddLanguage(ctx *gin.Context) {
	var language entity.ProgrammingLanguage
	err := ctx.ShouldBindJSON(&language)
	if err != nil {
		ctx.JSON(400, response.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "all field is required, please input all field",
		})
		return
	}

	err = helper.Validate.Struct(language)
	if err != nil {
		ctx.JSON(400, response.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "all field is required, please input all field",
		})
		return
	}

	result, err := controller.languageService.AddLanguage(language)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error{
			StatusCode: http.StatusBadRequest,
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

func (controller *LanguageControllers) GetAllLanguage(ctx *gin.Context) {
	result, err := controller.languageService.FetchAllLanguage()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error{
			StatusCode: http.StatusBadRequest,
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
