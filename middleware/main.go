package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func flow(h http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	appliedMiddleware := h
	for _, middleware := range middlewares {
		appliedMiddleware = middleware(appliedMiddleware)
	}
	return appliedMiddleware
}

func handleIndexRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello from middleware")
}

func requestLogger(l *log.Logger) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			l.Printf("started request to %s", r.URL)
			next(w, r)
			l.Printf("completed request to %s in %s", r.URL, time.Since(start))
		}
	}
}

func main() {
	http.HandleFunc("/", flow(
		handleIndexRoute,
		requestLogger(log.New(os.Stdout, "", 0)),
	))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
