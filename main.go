package main

import (
	"log"

	"github.com/ilhaamms/backend-live/api"
	"github.com/ilhaamms/backend-live/controller"
	"github.com/ilhaamms/backend-live/models/config"
	"github.com/ilhaamms/backend-live/repository"
	"github.com/ilhaamms/backend-live/service"
)

func main() {

	config, err := config.NewAppConfig()
	if err != nil {
		log.Fatal(err)
	}

	languageRepo := repository.NewLanguageRepository()

	goDeveloperService := service.NewGoDeveloperService(*config)
	languageService := service.NewLanguageService(languageRepo)

	goDeveloperController := controller.NewGoDeveloperController(goDeveloperService)
	languageController := controller.NewLanguageController(languageService)

	api := api.NewAPI(*config, goDeveloperController, languageController)

	api.Run()
}
