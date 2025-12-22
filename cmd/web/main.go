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
	"portfolio/internals/services"
	"portfolio/internals/utils"
	"portfolio/posts"
	"time"

	"github.com/joho/godotenv"
)

var data models.DataModel
var hosted_url string
var postService *services.PostService

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
	err = jsonParser.Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	data.AllSkills = utils.GetAllTechnologies(data)
	data.ProjectTechnologies = utils.GetProjectTechnologies(data.Projects)

	postService, err = services.NewPostService(posts.GetPostFS())
	if err != nil {
		log.Fatal("Error Initializing Post service")
	}

}

func main() {
	router := http.NewServeMux()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/", http.StripPrefix("/static/", middleware.StaticCache(fs)))

	router.Handle("/", middleware.PageCache(handelers.IndexHandeler(data, postService.Posts)))
	// TODO: Enable when blogs are ready
	// router.Handle("/blogs", middleware.PageCache(handelers.BlogHandeler(postService)))
	// router.Handle("/blogs/{slug}", middleware.PageCache(handelers.ReadingHandeler(postService)))

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
