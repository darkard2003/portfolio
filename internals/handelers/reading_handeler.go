package handelers

import (
	"log"
	"net/http"
	"portfolio/internals/services"
	"portfolio/web/view/pages/reading"
)

func ReadingHandeler(postService *services.PostService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug")
		blog, ok := postService.PostCache[slug]
		if !ok {
			http.Error(w, "Not found", http.StatusNotFound)
			log.Printf("Not found: %s", slug)
			return
		}
		screen := reading.ReadingScreen(blog)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := screen.Render(r.Context(), w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Error rendering page: %v", err)
		}
	}
}
