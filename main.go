package main

import (
	"github.com/ilhaamms/backend-live/api"
	"github.com/ilhaamms/backend-live/controller"
	"github.com/ilhaamms/backend-live/service"
)

func main() {

	goDeveloperService := service.NewGoDeveloperService()
	goDeveloperController := controller.NewGoDeveloperController(goDeveloperService)

	api := api.NewAPI(goDeveloperController)

	api.Run()
}
