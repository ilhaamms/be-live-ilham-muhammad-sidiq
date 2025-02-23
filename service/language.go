package service

import (
	"errors"
	"strconv"

	"github.com/ilhaamms/backend-live/models/entity"
	"github.com/ilhaamms/backend-live/repository"
)

type LanguageService interface {
	FecthLanguageByID(id int) (*entity.ProgrammingLanguage, error)
	AddLanguage(language entity.ProgrammingLanguage) (*entity.ProgrammingLanguage, error)
	FetchAllLanguage() ([]entity.ProgrammingLanguage, error)
	DeleteLanguageByID(id int) error
	UpdateLanguageByID(id int, language entity.ProgrammingLanguage) error
}

type LanguageServices struct {
	languageRepo repository.LanguageRepository
}

func NewLanguageService(languageRepo repository.LanguageRepository) LanguageService {
	return &LanguageServices{languageRepo: languageRepo}
}

func (s *LanguageServices) FecthLanguageByID(id int) (*entity.ProgrammingLanguage, error) {

	if id <= 0 {
		return nil, errors.New("parameter id must be greater than 0")
	}

	if id > len(entity.DataProgrammingLanguage) {
		return nil, errors.New("data languange with id " + strconv.Itoa(id) + " not found")
	}

	id = id - 1
	dataLanguage, err := s.languageRepo.FecthLanguageByID(id)
	if err != nil {
		return nil, err
	}

	return dataLanguage, nil
}

func (s *LanguageServices) AddLanguage(language entity.ProgrammingLanguage) (*entity.ProgrammingLanguage, error) {
	dataLanguage, err := s.languageRepo.AddLanguage(language)
	if err != nil {
		return nil, err
	}

	return &dataLanguage, nil
}

func (s *LanguageServices) FetchAllLanguage() ([]entity.ProgrammingLanguage, error) {
	dataLanguage, err := s.languageRepo.FetchAllLanguage()
	if err != nil {
		return nil, err
	}

	return dataLanguage, nil
}

func (s *LanguageServices) DeleteLanguageByID(id int) error {

	if id <= 0 {
		return errors.New("parameter id must be greater than 0")
	}

	if len(entity.DataProgrammingLanguage) == 0 {
		return errors.New("data language is empty")
	}

	if id > len(entity.DataProgrammingLanguage) {
		return errors.New("data languange with id " + strconv.Itoa(id) + " not found")
	}

	id = id - 1
	err := s.languageRepo.DeleteLanguageByID(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *LanguageServices) UpdateLanguageByID(id int, language entity.ProgrammingLanguage) error {
	if id <= 0 {
		return errors.New("parameter id must be greater than 0")
	}

	if id > len(entity.DataProgrammingLanguage) {
		return errors.New("data languange with id " + strconv.Itoa(id) + " not found")
	}

	id = id - 1
	err := s.languageRepo.UpdateLanguageByID(id, language)
	if err != nil {
		return err
	}

	return nil
}
