package service

import (
	"github.com/ilhaamms/backend-live/models/entity"
	"github.com/ilhaamms/backend-live/repository"
)

type LanguageService interface {
	FecthLanguageByID(id int) (*entity.ProgrammingLanguage, error)
}

type LanguageServices struct {
	languageRepo repository.LanguageRepository
}

func NewLanguageService(languageRepo repository.LanguageRepository) LanguageService {
	return &LanguageServices{languageRepo: languageRepo}
}

func (service *LanguageServices) FecthLanguageByID(id int) (*entity.ProgrammingLanguage, error) {
	dataLanguage, err := service.languageRepo.FecthLanguageByID(id)
	if err != nil {
		return nil, err
	}

	return dataLanguage, nil
}
