package handelers

import (
	"context"
	"log"
	"net/http"
	"portfolio/internals/models"
	"portfolio/web/view/pages/home"
	"portfolio/web/view/pages/notfound"
)

func IndexHandeler(data models.DataModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			notFound := notfound.NotFoundPage()
			if err := notFound.Render(context.Background(), w); err != nil {
				log.Printf("Error rendering 404 page: %v", err)
				http.Error(w, "404 Page Not Found", http.StatusNotFound)
			}
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		home := home.HomePage(data)

		if err := home.Render(context.Background(), w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Error rendering home page: %v", err)
		}
	}
}
