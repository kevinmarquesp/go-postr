package router

import (
	"log"
	"net/http"
	"time"
)

type wrappedWriter struct {
	http.ResponseWriter
	status int
}

func (ww *wrappedWriter) WriteHeader(status int) {
	ww.ResponseWriter.WriteHeader(status)

	ww.status = status
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		ww := &wrappedWriter{
			ResponseWriter: w,
			status:         http.StatusOK,
		}

		next.ServeHTTP(ww, r)

		log.Println("\t", time.Since(start), "\t", ww.status, r.Method, "\t", r.URL.Path)
	})
}
