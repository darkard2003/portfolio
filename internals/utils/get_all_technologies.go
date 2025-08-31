package utils

import (
	"portfolio/internals/models"
)

func GetAllTechnologies(data models.DataModel) []string {
	technologies := make([]string, 0)
	inclution_map := make(map[string]bool)

	for _, technology := range data.Skills.Languages {
		if !inclution_map[technology] {
			technologies = append(technologies, technology)
			inclution_map[technology] = true
		}
	}
	for _, technology := range data.Skills.Frameworks {
		if !inclution_map[technology] {
			technologies = append(technologies, technology)
			inclution_map[technology] = true
		}
	}
	for _, technology := range data.Skills.Databases {
		if !inclution_map[technology] {
			technologies = append(technologies, technology)
			inclution_map[technology] = true
		}
	}

	for _, project := range data.Projects {
		for _, technology := range project.Technologies {
			if !inclution_map[technology] {
				technologies = append(technologies, technology)
				inclution_map[technology] = true
			}
		}
	}

	return technologies
}
