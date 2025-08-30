package main

import (
	"fmt"
	"net/http"
	"os"
	"portfolio/internals/handelers"
)

func main() {
	router := http.NewServeMux()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Serve static files from the "static" directory
	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	router.HandleFunc("/", handelers.IndexHandeler)

	fmt.Println("Listening on port", port)
	http.ListenAndServe(":"+port, router)
}
