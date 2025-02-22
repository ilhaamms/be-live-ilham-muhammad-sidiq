package entity

type Relation struct {
	InfluencedBy []string `json:"influenced-by" validate:"required,min=1,dive,required"`
	Influences   []string `json:"influences" validate:"required,min=1,dive,required"`
}

type ProgrammingLanguage struct {
	Language       string   `json:"language" validate:"required"`
	Appeared       int      `json:"appeared" validate:"required"`
	Created        []string `json:"created" validate:"required,min=1,dive,required"`
	Functional     bool     `json:"functional" validate:"required"`
	ObjectOriented bool     `json:"object-oriented" validate:"required"`
	Relation       Relation `json:"relation" validate:"required"`
}

var DataProgrammingLanguage = []ProgrammingLanguage{
	{
		Language:       "C",
		Appeared:       1972,
		Created:        []string{"Dennis Ritchie"},
		Functional:     true,
		ObjectOriented: false,
		Relation: Relation{
			InfluencedBy: []string{"B", "ALGOL 68", "Assembly", "FORTRAN"},
			Influences:   []string{"C++", "Objective-C", "C#", "Java", "JavaScript", "PHP", "Go"},
		},
	},
}
