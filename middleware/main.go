package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"os"
)

type middleware func(http.HandlerFunc) http.HandlerFunc

func buildChain(f http.HandlerFunc, m ...middleware) http.HandlerFunc {
        // if our chain is done, use the original handlerfunc
        if len(m) == 0 {
                return f
        }
        // otherwise nest the handlerfuncs
        return m[0](buildChain(f, m[1:cap(m)]...))
}

func authMiddleware(f http.HandlerFunc) http.HandlerFunc {
	// one time scope setup area for middleware
	return func(w http.ResponseWriter, r *http.Request) {
		// ... pre handler functionality
		fmt.Println("start auth")
		f(w, r)
		fmt.Println("end auth")
		// ... post handler functionality
	}
}


func handleIndexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handling index route")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello from middleware")
}

func requestLogger(l *log.Logger) middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			fmt.Printf("started request to %s", r.URL)
			next(w, r)
			l.Printf("completed request to %s in %s", r.URL, time.Since(start))
		}
	}
}

func main() {
	http.HandleFunc("/", buildChain(
		handleIndexRoute,
		requestLogger(log.New(os.Stdout, "", 0)),
		authMiddleware,
	))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
