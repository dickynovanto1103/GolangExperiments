package main

import (
	"fmt"
	"log"
	"net/http"
)

type myHandler struct {
}

func NewMyHandler() myHandler {
	return myHandler{}
}

func (m myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello world, request method: %v", r.Method)
}

func main() {
	port := "8080"
	server := &http.Server{
		Addr:              fmt.Sprintf(":%v", port),
		Handler:           NewMyHandler(),
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
	}

	log.Printf("listening to port: %v", port)
	log.Fatal(server.ListenAndServe())
}
