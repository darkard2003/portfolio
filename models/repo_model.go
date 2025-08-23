package models

type Repo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	RepoUrl     string `json:"html_url"`
}
