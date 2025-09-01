package models

import "encoding/json"

type Project struct {
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Url          string   `json:"url"`
	Technologies []string `json:"technologies"`
}

func (p *Project) TechnologiesJson() string {
	jsonBytes, err := json.Marshal(p.Technologies)
	if err != nil {
		return "[]"
	}
	return string(jsonBytes)
}
