package main

import (
	"log"
	"net/http"
	"time"
)

func longRunningFunc(w http.ResponseWriter, r *http.Request) {
	log.Printf("request comes")
	time.Sleep(2 * time.Second)
	log.Printf("still writes hello world")
	w.Write([]byte("hello world"))
}

func main() {
	http.HandleFunc("/slow", longRunningFunc)

	log.Printf("start listening")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
