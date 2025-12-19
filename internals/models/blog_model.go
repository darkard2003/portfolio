package models

import (
	"html/template"
	"time"
)

type BlogModel struct {
	Title       string `yaml:"title"`
	Date        string `yaml:"date"`
	Slug        string `yaml:"slug"`
	Description string `yaml:"description"`
	Content     string
	ContentHTML  template.HTML
	DateObj     time.Time
}
