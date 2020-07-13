package main

import (
	"bytes"
	"encoding/json"
	"log"
)

type Student struct {
	Id int
	Name string
}


type TestStruct struct {
	Mapper map[string]int `json:"testmap"`
}

func createJSONString(ob interface{}) string {
	b, _ := json.Marshal(ob)
	return string(b)
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

	testStruct := &TestStruct{
		Mapper: make(map[string]int),
	}
	testStruct.Mapper["json"] = 1
	b, _ := json.Marshal(testStruct)
	log.Println("hasil marshall: ", string(b))

	ans := createJSONString(testStruct)
	log.Println("string: ", ans)

	name := "dicky"
	err = json.NewEncoder(bytes.NewBuffer([]byte(name))).Encode(struct{}{})
	log.Println("err encode: ", err)
}
