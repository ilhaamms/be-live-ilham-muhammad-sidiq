package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilhaamms/backend-live/helper"
	"github.com/ilhaamms/backend-live/service"
)

type PalindromController interface {
	GetPalindrom(ctx *gin.Context)
}

type PalindromControllers struct {
	palindromService service.PalindromService
}

func NewPalindromController(palindromService service.PalindromService) PalindromController {
	return &PalindromControllers{
		palindromService: palindromService,
	}
}

func (c *PalindromControllers) GetPalindrom(ctx *gin.Context) {

	text := ctx.Query("text")
	if text == "" {
		err := "parameter text is required, please input text"
		helper.ResponseJsonError(ctx, http.StatusBadRequest, err)
		return
	}

	result := c.palindromService.IsPalindrom(text)
	if result != "Palindrome" {
		err := "Not Palindrome"
		helper.ResponseJsonError(ctx, http.StatusBadRequest, err)
		return
	}

	message := "Palindrome"
	helper.ResponseJsonSuccess(ctx, http.StatusOK, message, nil)
}
