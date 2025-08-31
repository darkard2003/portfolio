package handelers

import (
	"context"
	"log"
	"net/http"
	"portfolio/web/view/pages/test"
)

func TestHandeler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	test := test.TestPage()
	if err := test.Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error rendering test page: %v", err)
	}

}
