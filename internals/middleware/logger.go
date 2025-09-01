package middleware

import (
	"log"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the ip also
		start := time.Now()
		rw := newResponseWriter(w)
		next.ServeHTTP(rw, r)
		end := time.Now()

		duration := end.Sub(start)

		log.Printf("%s %s FROM %s %d %s", r.Method, r.RequestURI, r.RemoteAddr, rw.statusCode, duration)
	})
}
