package models

type Portfolio struct {
	Name     string    `json:"name"`
	Contacts Contacts  `json:"contacts"`
	Summary  string    `json:"summary"`
	Projects []Project `json:"projects"`
	Skills   Skills    `json:"skills"`
	Hobbies  []string  `json:"hobbies"`
}

type Contacts struct {
	Email    string   `json:"email"`
	Titles   []string `json:"titles"`
	Phone    string   `json:"phone"`
	LinkedIn Social   `json:"linkedin"`
	Github   Social   `json:"github"`
}

type Social struct {
	Handle string `json:"handle"`
	URL    string `json:"url"`
}

type Project struct {
	Title        string   `json:"title"`
	Technologies []string `json:"technologies"`
	URL          string   `json:"url"`
	Description  string   `json:"description"`
}

type Skills struct {
	Languages  []string `json:"languages"`
	Frameworks []string `json:"frameworks"`
	Databases  []string `json:"databases"`
}
