package middleware

import (
	"net/http"
	"strings"
)

func CacheControlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, ".js") || strings.HasSuffix(r.URL.Path, ".css") {
			w.Header().Set("Cache-Control", "public, max-age=31536000") // 1 year
		}
		next.ServeHTTP(w, r)
	})
}
