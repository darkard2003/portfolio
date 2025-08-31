package handelers

import (
	"context"
	"log"
	"net/http"
	"portfolio/internals/models"
	"portfolio/web/view/pages/home"
)

func IndexHandeler(data models.DataModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		home := home.HomePage(data)

		if err := home.Render(context.Background(), w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Error rendering home page: %v", err)
		}
	}
}
