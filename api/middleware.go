package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs the incoming requests and their response time
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		log.Printf("Completed %s in %v", r.URL.Path, duration)
	})
}
