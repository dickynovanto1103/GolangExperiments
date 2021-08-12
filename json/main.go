package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Student struct {
	Id   int
	Name string
}

type TestStruct struct {
	Mapper map[string]int `json:"testmap"`
}

func createJSONString(ob interface{}) string {
	b, _ := json.Marshal(ob)
	return string(b)
}

type Member struct {
	Name    string   `json:"name"`
	AgeList []uint64 `json:"age_list"`
}

type Coba struct {
	Members []*Member `json:"members"`
}

func testUnmarshallNilValue() {
	var value *Coba
	err := json.Unmarshal(nil, value)
	log.Printf("err: %v value: %v", err, value)
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

	coba := &Coba{
		Members: []*Member{
			{
				Name:    "dicky",
				AgeList: []uint64{1, 2, 3},
			},
			{
				Name:    "asdf",
				AgeList: []uint64{1, 2, 4},
			},
		},
	}

	log.Printf("hasilMarshall: %v", createJSONString(coba))
	testUnmarshallNilValue()
	tryJsonEncoder()
	tryJsonDecoder()
}

func tryJsonEncoder() {
	log.Printf("trying json encoder")
	file, err := ioutil.TempFile("./", "dicky_test*.txt")
	if err != nil {
		log.Panicf("error in creating temp file, err: %v", err)
	}
	encoder := json.NewEncoder(file)
	student := []Student{
		{
			Id:   12,
			Name: "dicky",
		},
		{
			Id:   13,
			Name: "dicky13",
		},
	}

	if err := encoder.Encode(student); err != nil {
		log.Printf("error in encoding, err: %v", err)
	}
	os.Rename(file.Name(), "dicky_test.txt")
	log.Printf("finished encoding")
}

func tryJsonDecoder() {
	log.Printf("trying json decoder")
	file, err := os.Open("dicky_test.txt")
	if err != nil {
		log.Panicf("error in opening file, err: %v", err)
	}

	decoder := json.NewDecoder(file)
	students := &[]Student{}
	if err = decoder.Decode(students); err != nil {
		log.Panicf("error in decoding, err: %v", err)
	}
	log.Printf("student from decoding: %+v", students)
	for _, student := range *students {
		log.Printf("student: %+v", student)
	}
}
