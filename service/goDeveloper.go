package service

import (
	"github.com/ilhaamms/backend-live/models/config"
)

type GoDeveloperService interface {
	GoDeveloper() (string, error)
}

type GoDeveloperServices struct {
	config config.AppConfig
}

func NewGoDeveloperService(config config.AppConfig) GoDeveloperService {
	return &GoDeveloperServices{config: config}
}

func (service *GoDeveloperServices) GoDeveloper() (string, error) {
	return service.config.GolangDeveloper, nil
}
