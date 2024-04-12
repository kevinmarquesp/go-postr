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

// Reimplementation of the http.ResponseWriter.WriteHeader()[^1][^2] function,
// it will execute that function normally but will also store the response
// status on the router.wrappedWriter struct.
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

		// [^1]: Pass the custom writer object here.
		next.ServeHTTP(ww, r)

		// [^2]: Now have access to the status value right here.
		log.Println("\t", time.Since(start), "\t", ww.status, r.Method, "\t", r.URL.Path)
	})
}
