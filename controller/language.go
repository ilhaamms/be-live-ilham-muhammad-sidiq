package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ilhaamms/backend-live/models/entity"
	"github.com/ilhaamms/backend-live/models/response"
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

	if id <= 0 {
		ctx.JSON(400, response.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "parameter id must be greater than 0",
		})
		return
	}

	if id > len(entity.DataProgrammingLanguage) {
		ctx.JSON(400, response.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "parameter id must be less than or equal to " + strconv.Itoa(len(entity.DataProgrammingLanguage)),
		})
		return
	}

	id = id - 1
	result, err := controller.languageService.FecthLanguageByID(id)
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
