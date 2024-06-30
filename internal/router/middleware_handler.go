package router

import (
	"net/http"
	"regexp"
	"time"

	"github.com/charmbracelet/log"
)

const EXCLUDE_MIDDLEWARE_MATCH = `/favicon\.ico`

// Wrapper to hold the writer status when updating the status code.
type StatusWriter struct {
	http.ResponseWriter

	status int
}

func (sw *StatusWriter) WriteHeader(updatedStatus int) {
	sw.ResponseWriter.WriteHeader(updatedStatus)

	sw.status = updatedStatus
}

func MiddlewareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timerStart := time.Now()

		sw := &StatusWriter{
			ResponseWriter: w,
			status:         http.StatusOK,
		}

		if regexp.MustCompile(EXCLUDE_MIDDLEWARE_MATCH).MatchString(r.URL.Path) {
			handler.ServeHTTP(sw, r)
			return
		}

		if Middleware(sw, r) {
			handler.ServeHTTP(sw, r)
		}

		log.Printf("%s\t%d %s\t%s", time.Since(timerStart), sw.status, r.Method, r.URL.Path)
	})
}
