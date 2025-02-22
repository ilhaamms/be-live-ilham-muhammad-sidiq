package repository

import "github.com/ilhaamms/backend-live/models/entity"

type LanguageRepository interface {
	FecthLanguage() (entity.ProgrammingLanguage, error)
}

type languageRepository struct {
}

func NewLanguageRepository() LanguageRepository {
	return &languageRepository{}
}

func (repository *languageRepository) FecthLanguage() (entity.ProgrammingLanguage, error) {

	programmingLanguage := entity.ProgrammingLanguage{
		Language:       "C",
		Appeared:       1972,
		Created:        []string{"Dennis Ritchie"},
		Functional:     true,
		ObjectOriented: false,
		Relation: entity.Relation{
			InfluencedBy: []string{"B", "ALGOL 68", "Assembly", "FORTRAN"},
			Influences:   []string{"C++", "Objective-C", "C#", "Java", "JavaScript", "PHP", "Go"},
		},
	}

	return programmingLanguage, nil
}
