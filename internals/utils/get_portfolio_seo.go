package utils

import (
	"fmt"
	"portfolio/internals/models"
)

func GetPortfolioSEO(data models.DataModel) models.SEOModel {
	return models.NewSEOModel(
		fmt.Sprintf("%s's Portfolio", data.Name),
		data.Description,
		data.Url,
		"index, follow",
		"website",
		"/static/images/banner.png",
	)

}
