package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
	"sync"
)

var gzPool = sync.Pool{
	New: func() interface{} {
		return gzip.NewWriter(io.Discard)
	},
}

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
	gz           *gzip.Writer
	wroteHeader  bool
	shouldGzip   bool
	code         int
}

func (w *gzipResponseWriter) WriteHeader(code int) {
	if w.wroteHeader {
		return
	}
	w.code = code // Delay writing header until we determine if we gzip
	w.wroteHeader = true
	
	// Check content type to decide if we should gzip
	ct := w.Header().Get("Content-Type")
	if strings.Contains(ct, "text/") || 
	   strings.Contains(ct, "application/javascript") || 
	   strings.Contains(ct, "application/json") || 
	   strings.Contains(ct, "image/svg+xml") ||
	   strings.Contains(ct, "application/xml") {
		w.shouldGzip = true
	}
	
	if w.shouldGzip {
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Del("Content-Length")
	}
	
	w.ResponseWriter.WriteHeader(code)
}

func (w *gzipResponseWriter) Write(b []byte) (int, error) {
	if !w.wroteHeader {
		// If WriteHeader wasn't called explicitly, default to 200 and check content-type detection
		// However, standard lib often sniffs content type in WriteHeader if not set.
		// For safety, we can just call WriteHeader(200) which triggers our logic.
		w.WriteHeader(http.StatusOK)
	}

	if w.shouldGzip {
		if w.gz == nil {
			gz := gzPool.Get().(*gzip.Writer)
			gz.Reset(w.ResponseWriter)
			w.gz = gz
			w.Writer = gz
		}
		return w.gz.Write(b)
	}
	return w.ResponseWriter.Write(b)
}

func (w *gzipResponseWriter) Flush() {
	if w.gz != nil {
		w.gz.Flush()
	}
	if f, ok := w.ResponseWriter.(http.Flusher); ok {
		f.Flush()
	}
}

func (w *gzipResponseWriter) Close() error {
	if w.gz != nil {
		err := w.gz.Close()
		gzPool.Put(w.gz)
		return err
	}
	return nil
}

func GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Vary", "Accept-Encoding")
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		gzw := &gzipResponseWriter{
			ResponseWriter: w,
			code:           http.StatusOK,
		}
		defer gzw.Close()

		next.ServeHTTP(gzw, r)
	})
}
