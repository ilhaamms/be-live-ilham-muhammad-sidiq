package service

import (
	"errors"
	"strconv"

	"github.com/ilhaamms/backend-live/models/entity"
	"github.com/ilhaamms/backend-live/repository"
)

type LanguageService interface {
	FecthLanguageByID(id int) (*entity.ProgrammingLanguage, error)
	AddLanguage(language entity.ProgrammingLanguage) ([]entity.ProgrammingLanguage, error)
	FetchAllLanguage() ([]entity.ProgrammingLanguage, error)
	DeleteLanguageByID(id int) error
}

type LanguageServices struct {
	languageRepo repository.LanguageRepository
}

func NewLanguageService(languageRepo repository.LanguageRepository) LanguageService {
	return &LanguageServices{languageRepo: languageRepo}
}

func (service *LanguageServices) FecthLanguageByID(id int) (*entity.ProgrammingLanguage, error) {

	if id <= 0 {
		return nil, errors.New("parameter id must be greater than 0")
	}

	if id > len(entity.DataProgrammingLanguage) {
		return nil, errors.New("parameter id must be less than or equal to " + strconv.Itoa(len(entity.DataProgrammingLanguage)))
	}

	id = id - 1
	dataLanguage, err := service.languageRepo.FecthLanguageByID(id)
	if err != nil {
		return nil, err
	}

	return dataLanguage, nil
}

func (service *LanguageServices) AddLanguage(language entity.ProgrammingLanguage) ([]entity.ProgrammingLanguage, error) {
	dataLanguage, err := service.languageRepo.AddLanguage(language)
	if err != nil {
		return nil, err
	}

	return dataLanguage, nil
}

func (service *LanguageServices) FetchAllLanguage() ([]entity.ProgrammingLanguage, error) {
	dataLanguage, err := service.languageRepo.FetchAllLanguage()
	if err != nil {
		return nil, err
	}

	return dataLanguage, nil
}

func (service *LanguageServices) DeleteLanguageByID(id int) error {
	if len(entity.DataProgrammingLanguage) == 0 {
		return errors.New("data language is empty")
	}

	if id <= 0 {
		return errors.New("parameter id must be greater than 0")
	}

	if id > len(entity.DataProgrammingLanguage) {
		return errors.New("parameter id must be less than or equal to " + strconv.Itoa(len(entity.DataProgrammingLanguage)))
	}

	id = id - 1
	err := service.languageRepo.DeleteLanguageByID(id)
	if err != nil {
		return err
	}

	return nil
}
