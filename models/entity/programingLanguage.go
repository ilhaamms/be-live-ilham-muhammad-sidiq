package entity

type Relation struct {
	InfluencedBy []string `json:"influenced-by"`
	Influences   []string `json:"influences"`
}

type ProgrammingLanguage struct {
	Language       string   `json:"language"`
	Appeared       int      `json:"appeared"`
	Created        []string `json:"created"`
	Functional     bool     `json:"functional"`
	ObjectOriented bool     `json:"object-oriented"`
	Relation       Relation `json:"relation"`
}
