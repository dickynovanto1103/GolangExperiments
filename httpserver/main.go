package main

import (
	"encoding/json"
	"github.com/dickynovanto1103/GolangExperiments/jsonprotobuf/student"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method{
	case "GET":
		w.Write([]byte("hello"))
	case "POST":
		decoder := json.NewDecoder(r.Body)
		student := &student.Student{}
		if err := decoder.Decode(&student); err != nil {
			log.Panicf("error decoding student, err: %v", err)
		}
		log.Printf("r: %v body: %v form: %v", r, r.Body, r.Form)
		log.Println("student: ", student)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("POST request successful, student: " + student.String()))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8181", nil))
}
