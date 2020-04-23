package main

import (
	"bytes"
	"encoding/json"
	"github.com/dickynovanto1103/GolangExperiments/jsonprotobuf/model"
	"github.com/dickynovanto1103/GolangExperiments/jsonprotobuf/student"
	"github.com/gogo/protobuf/jsonpb"
	"io/ioutil"
	"log"
)

func main() {
	file, err := ioutil.ReadFile("jsonprotobuf/data/data.json")
	if err != nil {
		panic("error reading file " + err.Error())
	}
	data := &model.Student{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Println("error unmarshal data, err: ", err)
		return
	}
	log.Println("data: ", data)

	studentPb := &student.Student{}

	err = jsonpb.Unmarshal(bytes.NewBuffer(file), studentPb)
	if err != nil {
		log.Println("error unmarshall studentpb, err:", err)
		return
	}

	log.Printf("studentPb: %v\n", studentPb)
	marshaller := &jsonpb.Marshaler{}
	newBuffer := bytes.NewBuffer([]byte{})
	err = marshaller.Marshal(newBuffer, studentPb)
	if err != nil {
		log.Println("Err: ", err)
		return
	}
	ans := newBuffer.Bytes()
	log.Println("ans: ", string(ans))
}
