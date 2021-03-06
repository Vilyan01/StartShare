package main

import (
	"log"
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf(
			"%10s\t%10s\t%10s\t%10s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
