package main

import (
	"log"
	"net/http"
	"time"
)

func GenericMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO replace with zap/fuller logging middleware
		start := time.Now()
		log.Println(r.RequestURI)
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
		log.Printf(
			"%s\t%s\tTook: %s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}
