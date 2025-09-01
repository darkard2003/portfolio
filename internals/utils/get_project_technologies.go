package utils

import "portfolio/internals/models"

func GetProjectTechnologies(projects []models.Project) []string {
	allTechnologies := make([]string, 0)
	shadow_map := make(map[string]bool)

	for _, project := range projects {
		for _, tech := range project.Technologies {
			if _, ok := shadow_map[tech]; !ok {
				allTechnologies = append(allTechnologies, tech)
				shadow_map[tech] = true
			}
		}
	}

	return allTechnologies
}
