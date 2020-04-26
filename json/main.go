package main

import (
	"encoding/json"
	"log"
)

type Student struct {
	Id int
	Name string
}

//if we don't set the name of the fields in json format, for example, Id int `json:"id"`, then the fields will just become the original fields name
func main() {
	student := &Student{}
	student.Name = "diky"
	student.Id = 123
	log.Println("student: ", student)
	byteArr, err := json.Marshal(student)
	if err != nil {
		log.Println("err: ", err)
		return
	}
	log.Println("byteArr: ", string(byteArr))
}
