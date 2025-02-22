package repository

import "github.com/ilhaamms/backend-live/models/entity"

type LanguageRepository interface {
	FecthLanguageByID(id int) (*entity.ProgrammingLanguage, error)
	AddLanguage(language entity.ProgrammingLanguage) ([]entity.ProgrammingLanguage, error)
	FetchAllLanguage() ([]entity.ProgrammingLanguage, error)
	DeleteLanguageByID(id int) error
}

type languageRepository struct {
}

func NewLanguageRepository() LanguageRepository {
	return &languageRepository{}
}

func (repository *languageRepository) FecthLanguageByID(id int) (*entity.ProgrammingLanguage, error) {
	dataProgramingLanguage := entity.DataProgrammingLanguage[id]
	return &dataProgramingLanguage, nil
}

func (repository *languageRepository) AddLanguage(language entity.ProgrammingLanguage) ([]entity.ProgrammingLanguage, error) {
	entity.DataProgrammingLanguage = append(entity.DataProgrammingLanguage, language)
	return entity.DataProgrammingLanguage, nil
}

func (repository *languageRepository) FetchAllLanguage() ([]entity.ProgrammingLanguage, error) {
	return entity.DataProgrammingLanguage, nil
}

func (repository *languageRepository) DeleteLanguageByID(id int) error {
	entity.DataProgrammingLanguage = append(entity.DataProgrammingLanguage[:id], entity.DataProgrammingLanguage[id+1:]...)
	return nil
}
