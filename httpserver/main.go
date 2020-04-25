package main

import (
	"encoding/json"
	"github.com/dickynovanto1103/GolangExperiments/jsonprotobuf/student"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Panicf("error in parsing form, err: %v", err)
	}
	decoder := json.NewDecoder(r.Body)
	student := &student.Student{}
	if err := decoder.Decode(&student); err != nil {
		log.Panicf("error decoding student, err: %v", err)
	}
	log.Printf("r: %v body: %v form: %v", r, r.Body, r.Form)
	log.Println("student: ", student)
	w.Write([]byte("hello"))
}

func main() {
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8181", nil))
}
