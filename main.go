package main

import (
	"embed"
	"encoding/json"
	"html/template"
	"io/fs"
	"log"
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

func main() {
	staticFS, err := fs.Sub(static, "static")
	if err != nil {
		log.Fatal(err)
	}
	templateFS, err := fs.Sub(templates, "templates")
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.StaticFS("/static", http.FS(staticFS))
	templ := template.Must(template.ParseFS(templateFS, "*.html"))
	router.SetHTMLTemplate(templ)

	var portfolio models.Portfolio
	err = json.Unmarshal(data, &portfolio)
	if err != nil {
		panic(err)
	}

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", portfolio)
	})

	// Use PORT environment variable or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)

}
