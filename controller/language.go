package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ilhaamms/backend-live/helper"
	"github.com/ilhaamms/backend-live/models/entity"
	"github.com/ilhaamms/backend-live/models/response"
	"github.com/ilhaamms/backend-live/service"
)

type LanguageController interface {
	GetLanguageByID(ctx *gin.Context)
	AddLanguage(ctx *gin.Context)
	GetAllLanguage(ctx *gin.Context)
	RemoveLanguageByID(ctx *gin.Context)
	ChangeLanguageByID(ctx *gin.Context)
}

type LanguageControllers struct {
	languageService service.LanguageService
}

func NewLanguageController(languageService service.LanguageService) LanguageController {
	return &LanguageControllers{
		languageService: languageService,
	}
}

func (c *LanguageControllers) GetLanguageByID(ctx *gin.Context) {

	idParam := ctx.Param("id")
	if idParam == "" {
		err := "parameter id is required, please input id"
		helper.ResponseJsonError(ctx, http.StatusBadRequest, err)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		err := "parameter id must be a number"
		helper.ResponseJsonError(ctx, http.StatusBadRequest, err)
		return
	}

	result, err := c.languageService.FecthLanguageByID(id)
	if err != nil {

		if strings.Contains(err.Error(), "not found") {
			helper.ResponseJsonError(ctx, http.StatusNotFound, err.Error())
		} else {
			helper.ResponseJsonError(ctx, http.StatusBadRequest, err.Error())
		}

		return
	}

	success := "Success get language by id " + idParam
	helper.ResponseJsonSuccess(ctx, http.StatusOK, success, result)
}

func (c *LanguageControllers) AddLanguage(ctx *gin.Context) {
	var language entity.ProgrammingLanguage
	err := ctx.ShouldBindJSON(&language)
	if err != nil {
		err := "all field is required, please input all field"
		helper.ResponseJsonError(ctx, http.StatusBadRequest, err)
		return
	}

	err = helper.Validate.Struct(language)
	if err != nil {
		err := "all field is required, please input all field"
		helper.ResponseJsonError(ctx, http.StatusBadRequest, err)
		return
	}

	result, err := c.languageService.AddLanguage(language)
	if err != nil {
		helper.ResponseJsonError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	success := "Success add language"
	helper.ResponseJsonSuccess(ctx, http.StatusCreated, success, result)
}

func (c *LanguageControllers) GetAllLanguage(ctx *gin.Context) {

	result, err := c.languageService.FetchAllLanguage()
	if err != nil {
		if strings.Contains(err.Error(), "page not found") {
			helper.ResponseJsonError(ctx, http.StatusNotFound, err.Error())
		} else {
			helper.ResponseJsonError(ctx, http.StatusBadRequest, err.Error())
		}

		return
	}

	ctx.JSON(http.StatusOK, response.Success{
		StatusCode: http.StatusOK,
		Message:    "Success get all language",
		Data:       result,
	})
}

func (c *LanguageControllers) RemoveLanguageByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	if idParam == "" {
		err := "parameter id is required, please input id"
		helper.ResponseJsonError(ctx, http.StatusBadRequest, err)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		err := "parameter id must be a number"
		helper.ResponseJsonError(ctx, http.StatusBadRequest, err)
		return
	}

	err = c.languageService.DeleteLanguageByID(id)
	if err != nil {
		helper.ResponseJsonError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	success := "Success delete language by id " + idParam
	helper.ResponseJsonSuccess(ctx, http.StatusOK, success, nil)
}

func (c *LanguageControllers) ChangeLanguageByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	if idParam == "" {
		err := "parameter id is required, please input id"
		helper.ResponseJsonError(ctx, http.StatusBadRequest, err)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		err := "parameter id must be a number"
		helper.ResponseJsonError(ctx, http.StatusBadRequest, err)
		return
	}

	var language entity.ProgrammingLanguage
	err = ctx.ShouldBindJSON(&language)
	if err != nil {
		err := "all field is required, please input all field"
		helper.ResponseJsonError(ctx, http.StatusBadRequest, err)
		return
	}

	err = c.languageService.UpdateLanguageByID(id, language)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			helper.ResponseJsonError(ctx, http.StatusNotFound, err.Error())
		} else {
			helper.ResponseJsonError(ctx, http.StatusBadRequest, err.Error())
		}

		return
	}

	success := "Success update language by id " + idParam
	helper.ResponseJsonSuccess(ctx, http.StatusOK, success, nil)
}
