package service

import (
	"github.com/ilhaamms/backend-live/models/entity"
)

type GoDeveloperService interface {
	GoDeveloper() (string, error)
}

type GoDeveloperServices struct {
}

func NewGoDeveloperService() GoDeveloperService {
	return &GoDeveloperServices{}
}

func (s *GoDeveloperServices) GoDeveloper() (string, error) {

	goDeveloper := entity.Developer{
		GoDeveloper: "Hello Go developers",
	}

	return goDeveloper.GoDeveloper, nil
}
