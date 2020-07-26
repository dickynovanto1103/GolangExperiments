package main

import (
	"golang.org/x/time/rate"
	"log"
	"net/http"
)

var limiter = rate.NewLimiter(1, 3)

type handler struct {

}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello world"))
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("im fine"))
}

func limit(defaultHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		defaultHandler.ServeHTTP(w, r)
	})
}

func main() {
	//newHandler := NewHandler()
	mux := http.NewServeMux()
	mux.HandleFunc("/", okHandler)

	log.Println("listening on 8080")
	if err := http.ListenAndServe(":8080", limit(mux)); err != nil {
		log.Fatalf("error: %v", err)
	}
}