package middleware

import (
	"net/http"
)

// CacheMiddleware sets Cache-Control headers for static assets
func CacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set generic cache control for static files (e.g., 1 year)
		w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
		next.ServeHTTP(w, r)
	})
}
