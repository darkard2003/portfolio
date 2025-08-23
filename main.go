package main

import (
	"embed"
	"encoding/json"
	"html/template"
	"io/fs"
	"net/http"
	"os"
	"portfolio/models"

	"github.com/gin-gonic/gin"
)

//go:embed static/*
var static embed.FS

//go:embed templates/*
var templates embed.FS

//go:embed data.json
var data []byte

// Filesystems
var staticFS fs.FS
var templateFS fs.FS

// portfolio models
var portfolio models.Portfolio

func init() {
	var err error
	staticFS, err = fs.Sub(static, "static")
	if err != nil {
		panic(err)
	}

	templateFS, err = fs.Sub(templates, "templates")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &portfolio)
	if err != nil {
		panic(err)
	}
}

func main() {
	templ := template.Must(template.ParseFS(templateFS, "*.html"))

	router := gin.Default()
	router.SetHTMLTemplate(templ)
	router.StaticFS("/static", http.FS(staticFS))

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", portfolio)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)

}
