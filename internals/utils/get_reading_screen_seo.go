package utils

import "portfolio/internals/models"

func GetReadingScreenSEO(blog models.PostModel, url string) models.SEOModel {
	return models.NewSEOModel(
		blog.Title,
		blog.Description,
		url,
		"index, follow, max-image-preview:large",
		"article",
		// TODO: add image later when added in post model
		"",
	)

}
