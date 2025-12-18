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

	"github.com/joho/godotenv"
)

var data models.DataModel
var hosted_url string

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default values")
	}

	hosted_url = os.Getenv("HOSTED_URL")
	if hosted_url == "" {
		hosted_url = "http://localhost:8080"
	}

	log.Printf("Allowed CORS origin: %s\n", hosted_url)

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
	// Apply CacheMiddleware specifically to static files
	router.Handle("/static/", http.StripPrefix("/static/", middleware.CacheMiddleware(fs)))

	router.HandleFunc("/", handelers.IndexHandeler(data))

	server := &http.Server{
		Addr: ":" + port,
		Handler: middleware.Chain(router,
			middleware.Logger,
			middleware.GzipMiddleware, // Apply Gzip globally
			func(h http.Handler) http.Handler {
				return middleware.CORSMiddleware(h, hosted_url)
			},
		),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	fmt.Printf("Server starting on port %s\n", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", port, err)
	}
}
