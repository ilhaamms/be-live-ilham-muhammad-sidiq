package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ilhaamms/backend-live/controller"
	"github.com/ilhaamms/backend-live/models/config"
)

type API struct {
	config                config.AppConfig
	goDeveloperController controller.GoDeveloperController
	languageController    controller.LanguageController
}

func NewAPI(
	config config.AppConfig,
	goDeveloperController controller.GoDeveloperController,
	languageController controller.LanguageController,
) *API {
	return &API{
		config:                config,
		goDeveloperController: goDeveloperController,
		languageController:    languageController,
	}
}

func (a *API) RegisterRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/", a.goDeveloperController.GoDeveloper)
	r.GET("/language", a.languageController.GetLanguage)

	// auth := r.Group("/auth")
	// {
	// 	auth.POST("/register", a.userController.Register)
	// 	auth.POST("/login", a.userController.Login)
	// }

	// r.POST("/authors", middleware.Auth(), a.authorController.CreateAuthor)
	// r.GET("/authors", middleware.Auth(), a.authorController.GetAllAuthor)
	// r.GET("/authors/:id", middleware.Auth(), a.authorController.GetAuthorsById)
	// r.DELETE("/authors/:id", middleware.Auth(), a.authorController.DeleteAuthorsById)
	// r.PUT("/authors/:id", middleware.Auth(), a.authorController.UpdateAuthorsById)

	// r.POST("/books", middleware.Auth(), a.bookController.CreateBook)
	// r.GET("/books", middleware.Auth(), a.bookController.GetAllBook)
	// r.GET("/books/:id", middleware.Auth(), a.bookController.GetBookById)
	// r.DELETE("/books/:id", middleware.Auth(), a.bookController.DeleteBookById)
	// r.PUT("/books/:id", middleware.Auth(), a.bookController.Update)

	return r
}

func (a *API) Run() {
	r := a.RegisterRoutes()

	log.Printf("Server running on port %s", a.config.PortService)
	r.Run(":" + a.config.PortService)
}
