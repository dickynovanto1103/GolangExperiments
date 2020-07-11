package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]string)
	m["dicky"] = "novanto"
	m["andre"] = "dwiyanto"
	w.WriteHeader(http.StatusInternalServerError)
	err := json.NewEncoder(w).Encode(m)
	log.Println("err encode:", err)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
