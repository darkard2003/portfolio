package main

import (
	"encoding/json"
	"net/http"
	"os"
	"portfolio/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	// Read and parse the JSON data
	file, err := os.ReadFile("data.json")
	if err != nil {
		panic(err)
	}

	var portfolio models.Portfolio
	err = json.Unmarshal(file, &portfolio)
	if err != nil {
		panic(err)
	}

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", portfolio)
	})

	router.Run(":8080")
}