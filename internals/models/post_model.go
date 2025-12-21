package models

import (
	"html/template"
	"time"
)

type PostModel struct {
	Title       string `yaml:"title"`
	Date        string `yaml:"date"`
	Description string `yaml:"description"`
	Author      string `yaml:"author"`
	Slug        string
	Content     string
	ContentHTML template.HTML
	DateObj     time.Time
}
