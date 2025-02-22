package repository

import "github.com/ilhaamms/backend-live/models/entity"

type LanguageRepository interface {
	FecthLanguageByID(id int) (*entity.ProgrammingLanguage, error)
	// AddLanguage(language entity.ProgrammingLanguage) (entity.ProgrammingLanguage, error)
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

// func (repository *languageRepository) AddLanguage(dataProgramingLanguage entity.ProgrammingLanguage) (entity.ProgrammingLanguage, error) {
// 	newDataProgramingLanguage := dataProgramingLanguage
// 	entity.DataProgrammingLanguage = newDataProgramingLanguage
// }
