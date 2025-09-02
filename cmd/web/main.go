package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"portfolio/internals/handelers"
	"portfolio/internals/middleware"
	"portfolio/internals/models"
	"portfolio/internals/utils"
	"time"
)

var data models.DataModel
var hosted_url string

func init() {

	hosted_url = os.Getenv("HOSTED_URL")
	if hosted_url == "" {
		hosted_url = "http://localhost:8080"
	}

	file, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}

	jsonParser := json.NewDecoder(file)
	jsonParser.Decode(&data)

	data.AllSkills = utils.GetAllTechnologies(data)
	data.ProjectTechnologies = utils.GetProjectTechnologies(data.Projects)
}

func main() {
	router := http.NewServeMux()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/", middleware.Chain(
		http.StripPrefix("/static/", fs),
		middleware.Logger,
		func(h http.Handler) http.Handler {
			return middleware.CORSMiddleware(h, hosted_url)
		},
	))

	router.HandleFunc("/", handelers.IndexHandeler(data))
	router.HandleFunc("/test", handelers.TestHandeler)

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      middleware.Logger(router),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	fmt.Printf("Server starting on port %s\n", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", port, err)
	}
}
