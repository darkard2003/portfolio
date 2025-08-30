package handelers

import (
	"context"
	"log"
	"net/http"
	"portfolio/web/view/pages"
)

func IndexHandeler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	home := pages.HomePage()

	if err := home.Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error rendering home page: %v", err)
	}
}
