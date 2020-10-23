package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"reflect"

	"github.com/dickynovanto1103/GolangExperiments/jsonprotobuf/model"
	"github.com/dickynovanto1103/GolangExperiments/jsonprotobuf/student"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

// experiment using jsonpb
func exp1(file []byte) {
	data := &model.Student{}
	err := json.Unmarshal(file, &data)
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

//exp2 directly use json.Unmarshall instead of json
func exp2(file []byte) {
	log.Println("exp2")
	data := &student.Student{}
	err := json.Unmarshal(file, &data)
	if err != nil {
		log.Println("error unmarshal data, err: ", err)
		return
	}

	log.Println("data: ", data, reflect.TypeOf(data))
	byteArr, err := json.Marshal(data)
	if err != nil {
		log.Println("err marshal json, err: ", err)
		return
	}
	log.Println("byteArr marshal result:", string(byteArr))
}

func exp3() {
	log.Printf("exp3")
	var bytes []byte

	unmarshalledStudent := &student.Student{}
	err := proto.Unmarshal(bytes, unmarshalledStudent)
	if err != nil {
		log.Printf("err: %v", err)
	}
	log.Printf("unmarshalledStudent: %v %T", unmarshalledStudent, unmarshalledStudent)
}

func main() {
	file, err := ioutil.ReadFile("jsonprotobuf/data/data.json")
	if err != nil {
		panic("error reading file " + err.Error())
	}
	exp1(file)
	exp2(file)
	exp3()
}
