package repository

import "github.com/ilhaamms/backend-live/models/entity"

type LanguageRepository interface {
	FecthLanguageByID(id int) (*entity.ProgrammingLanguage, error)
	AddLanguage(language entity.ProgrammingLanguage) (entity.ProgrammingLanguage, error)
	FetchAllLanguage() ([]entity.ProgrammingLanguage, error)
	DeleteLanguageByID(id int) error
	UpdateLanguageByID(id int, language entity.ProgrammingLanguage) error
}

type languageRepository struct {
}

func NewLanguageRepository() LanguageRepository {
	return &languageRepository{}
}

func (r *languageRepository) FecthLanguageByID(id int) (*entity.ProgrammingLanguage, error) {
	dataProgramingLanguage := entity.DataProgrammingLanguage[id]
	return &dataProgramingLanguage, nil
}

func (r *languageRepository) AddLanguage(language entity.ProgrammingLanguage) (entity.ProgrammingLanguage, error) {
	entity.DataProgrammingLanguage = append(entity.DataProgrammingLanguage, language)
	return language, nil
}

func (r *languageRepository) FetchAllLanguage() ([]entity.ProgrammingLanguage, error) {
	return entity.DataProgrammingLanguage, nil
}

func (r *languageRepository) DeleteLanguageByID(id int) error {
	entity.DataProgrammingLanguage = append(entity.DataProgrammingLanguage[:id], entity.DataProgrammingLanguage[id+1:]...)
	return nil
}

func (r *languageRepository) UpdateLanguageByID(id int, language entity.ProgrammingLanguage) error {

	dataLanguage := r.CheckUpdateDataLanguage(entity.DataProgrammingLanguage[id], language)

	entity.DataProgrammingLanguage[id] = dataLanguage
	return nil
}

func (r *languageRepository) CheckUpdateDataLanguage(oldDataLanguage, newDataLanguage entity.ProgrammingLanguage) entity.ProgrammingLanguage {
	if newDataLanguage.Language != "" {
		oldDataLanguage.Language = newDataLanguage.Language
	}

	if newDataLanguage.Appeared != 0 {
		oldDataLanguage.Appeared = newDataLanguage.Appeared
	}

	if len(newDataLanguage.Created) != 0 {
		oldDataLanguage.Created = newDataLanguage.Created
	}

	if newDataLanguage.Functional != nil {
		oldDataLanguage.Functional = newDataLanguage.Functional
	}

	if newDataLanguage.ObjectOriented != nil {
		oldDataLanguage.ObjectOriented = newDataLanguage.ObjectOriented
	}

	if newDataLanguage.Relation.InfluencedBy != nil {
		oldDataLanguage.Relation.InfluencedBy = newDataLanguage.Relation.InfluencedBy
	}

	if newDataLanguage.Relation.Influences != nil {
		oldDataLanguage.Relation.Influences = newDataLanguage.Relation.Influences
	}

	return oldDataLanguage
}
