package main

import (
	"log"
	"net/http"
	"time"
)

type Handler struct {

}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(500*time.Millisecond) //it reach timeout when writing a response is taking more than then "WriteTimeout"
	w.Write([]byte("dicky"))
}

func main() {
	server := &http.Server{
		Addr:              ":8080",
		Handler: &Handler{},
		ReadTimeout:       time.Second,
		WriteTimeout:      time.Millisecond,
		IdleTimeout:       time.Second,
		MaxHeaderBytes:    12,
	}
	err := server.ListenAndServe()
	log.Println("err: ", err)
}