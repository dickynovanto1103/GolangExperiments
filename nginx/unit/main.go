package main

import (
	"net/http"
	"unit.nginx.org/go"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func main() {
	http.HandleFunc("/", handler)

	unit.ListenAndServe(":8080", nil)
}
