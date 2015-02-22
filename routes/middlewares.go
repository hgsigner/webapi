package routes

import (
	"log"
	"net/http"
	"time"
)

func logMiddleware(mid http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf(">> %s %s", r.Method, r.URL.Path)
		mid.ServeHTTP(w, r)
		log.Printf("<< %s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}
