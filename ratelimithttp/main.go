package main

import (
	"log"
	"net/http"

	"golang.org/x/time/rate"
)

func main() {
	limiter := rate.NewLimiter(1, 3)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if !limiter.Allow() {
			writer.WriteHeader(http.StatusTooManyRequests)
			writer.Write([]byte("too many request"))
			return
		}

		log.Println("receiving http request")
		writer.Write([]byte("ok"))
	})

	http.ListenAndServe(":8080", nil)
}
