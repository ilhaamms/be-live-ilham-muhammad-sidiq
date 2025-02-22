package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilhaamms/backend-live/models/response"
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

func (controller *PalindromControllers) GetPalindrom(ctx *gin.Context) {

	text := ctx.Query("text")
	if text == "" {
		ctx.JSON(400, response.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "parameter text is required, please input text",
		})
		return
	}

	result := controller.palindromService.IsPalindrom(text)
	if result != "Palindrome" {
		ctx.JSON(http.StatusBadRequest, response.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "Not palindrome",
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Success{
		StatusCode: http.StatusOK,
		Message:    "Palindrome",
	})
}
