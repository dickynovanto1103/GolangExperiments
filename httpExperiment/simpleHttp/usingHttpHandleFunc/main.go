package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlerFunc1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to handler 1: %v, request url: %+v",
		r.FormValue("name"), r.URL)
}

func handlerFunc2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to handler 2, request method: %v", r.Method)
}

func main() {
	http.HandleFunc("/handle1", handlerFunc1)
	http.HandleFunc("/handle2", handlerFunc2)

	log.Printf("start listening")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Panicf("error in listen and serve, err: %v", err)
	}
}
