package models

type SEOModel struct {
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Url         string       `json:"url"`
	Robots      string       `json:"robots"`
	OG          OGModel      `json:"og"`
	Twitter     TwitterModel `json:"twitter"`
}

func NewSEOModel(
	title,
	description,
	url,
	robots,
	ogtype,
	image string,
) SEOModel {
	ogModel := NewOGModel(ogtype, url, title, description, image)
	card_type := "summary"
	if image != "" {
		card_type = "summary_large_image"
	}

	twitterModel := NewTwitterModel(card_type, title, description, image)
	return SEOModel{
		Title:       title,
		Description: description,
		Url:         url,
		Robots:      robots,
		OG:          ogModel,
		Twitter:     twitterModel,
	}
}

type OGModel struct {
	Type        string `json:"type"`
	Url         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func NewOGModel(ogType, url, title, description, image string) OGModel {
	return OGModel{
		Type:        ogType,
		Url:         url,
		Title:       title,
		Description: description,
		Image:       image,
	}

}

type TwitterModel struct {
	Card        string `json:"card"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func NewTwitterModel(card, title, description, image string) TwitterModel {
	return TwitterModel{
		Card:        card,
		Title:       title,
		Description: description,
		Image:       image,
	}
}
