package models

type DataModel struct {
	Name                string   `json:"name"`
	Titles              []string `json:"titles"`
	Contacts            Contacts `json:"contacts"`
	Summary             string   `json:"summary"`
	Projects            []Project
	ProjectTechnologies []string `json:"-"`
	Skills              Skills
	Hobbies             []string `json:"hobbies"`
	AllSkills           []string `json:"-"`
}
