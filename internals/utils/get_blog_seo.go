package utils

import "portfolio/internals/models"

func GetBlogSeo(data models.DataModel) models.SEOModel {
	return models.NewSEOModel(
		"All Blogs",                              // Title
		data.Summary,                             // Description
		data.Url+"/blogs",                        // URL
		"index, follow, max-image-preview:large", // Robots
		"website",                                // OG Type
		"/static/images/blog-banner.png",         // Image
	)
}
