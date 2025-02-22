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

func (s *GoDeveloperServices) GoDeveloper() (string, error) {
	return s.config.GolangDeveloper, nil
}
