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
		w.WriteHeader(http.StatusOK) //the write Header will only be set once, and will not be assigned to the other value later if set with the other values
		//w.WriteHeader(http.StatusInternalServerError) // if we set it twice, it will give warning of superfluous WriteHeader operation
		path := r.URL.String()
		log.Println("path: ", path)
		byteArr, _ := json.Marshal(student)

		w.Write([]byte("POST request successful, student: " + string(byteArr)))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handlerArray(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var arr []int32
	if err := decoder.Decode(&arr); err != nil {
		log.Println("error decoding, err: ", err)
	}
	log.Println("arr: ", arr)
	byteArr, _ := json.Marshal(arr)
	log.Println("byteArr: ", string(byteArr))
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/arr", handlerArray)

	log.Fatal(http.ListenAndServe(":8183", nil))
}
