package handelers

import (
	"context"
	"log"
	"net/http"
	"portfolio/internals/models"
	"portfolio/internals/services"
	"portfolio/internals/utils"
	"portfolio/web/view/pages/blog"
)

func BlogHandeler(postService *services.PostService, data models.DataModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		seo := utils.GetBlogSeo(data)
		blog := blog.Blog(postService.Posts, seo)
		if err := blog.Render(context.Background(), w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Error rendering page: %v", err)
		}
	}
}
