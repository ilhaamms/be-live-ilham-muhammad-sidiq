package service

import (
	"github.com/ilhaamms/backend-live/models/entity"
	"github.com/ilhaamms/backend-live/repository"
)

type LanguageService interface {
	FecthLanguage() (*entity.ProgrammingLanguage, error)
}

type LanguageServices struct {
	languageRepo repository.LanguageRepository
}

func NewLanguageService(languageRepo repository.LanguageRepository) LanguageService {
	return &LanguageServices{languageRepo: languageRepo}
}

func (service *LanguageServices) FecthLanguage() (*entity.ProgrammingLanguage, error) {
	dataLanguage, err := service.languageRepo.FecthLanguage()
	if err != nil {
		return nil, err
	}

	return &dataLanguage, nil
}
