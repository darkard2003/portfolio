package handelers

import (
	"context"
	"log"
	"net/http"
	"portfolio/internals/services"
	"portfolio/web/view/pages/blog"
)

func BlogHandeler(postService *services.PostService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		blog := blog.Blog(postService.Posts)

		if err := blog.Render(context.Background(), w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Error rendering page: %v", err)
		}
	}
}
