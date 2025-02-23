package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ilhaamms/backend-live/controller"
	"github.com/ilhaamms/backend-live/middleware"
	"github.com/ilhaamms/backend-live/models/config"
)

type API struct {
	config                config.AppConfig
	goDeveloperController controller.GoDeveloperController
	languageController    controller.LanguageController
	palindromController   controller.PalindromController
}

func NewAPI(
	config config.AppConfig,
	goDeveloperController controller.GoDeveloperController,
	languageController controller.LanguageController,
	palindromController controller.PalindromController,
) *API {
	return &API{
		config:                config,
		goDeveloperController: goDeveloperController,
		languageController:    languageController,
		palindromController:   palindromController,
	}
}

func (a *API) RegisterRoutes() *gin.Engine {
	r := gin.Default()

	r.NoRoute(middleware.MethodNotAllowedMiddleware())

	r.GET("/", a.goDeveloperController.GoDeveloper)
	r.GET("/palindrom", a.palindromController.GetPalindrom)

	language := r.Group("/language")
	{

		language.GET("/:id", a.languageController.GetLanguageByID)
		language.POST("/", a.languageController.AddLanguage)
		language.GET("/", a.languageController.GetAllLanguage)
		language.DELETE("/:id", a.languageController.RemoveLanguageByID)
		language.PATCH("/:id", a.languageController.ChangeLanguageByID)
	}

	return r
}

func (a *API) Run() {
	r := a.RegisterRoutes()

	log.Printf("Server running on port %s", a.config.PortService)
	r.Run(":" + a.config.PortService)
}
