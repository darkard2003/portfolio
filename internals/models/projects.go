package models

type Project struct {
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Url          string   `json:"url"`
	Technologies []string `json:"technologies"`
}
