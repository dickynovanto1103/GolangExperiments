package main

import (
	"log"
	"testing"

	student2 "github.com/dickynovanto1103/GolangExperiments/protobufRedis/student"
	"github.com/gogo/protobuf/proto"
)

func TestIDCardProtoSize(t *testing.T) {
	idCard := &student2.IDCard{
		Id: nil,
	}
	log.Printf("size XXXIdCard: %v size idCard: %v", idCard.XXX_Size(), proto.Size(idCard))
}

func TestStudentProtoSize(t *testing.T) {
	student := &student2.Student{
		Id:     proto.Int32(123),
		Name:   proto.String("aaaaaaaaaa"),
		IdCard: nil,
		Books: []*student2.Book{
			{
				Id:   nil,
				Name: proto.String("aa"),
			},
		},
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	marshalledStudent, _ := proto.Marshal(student)
	log.Printf("student size: %v size from proto: %v", student.XXX_Size(), proto.Size(student))
	log.Printf("byteLength of marshalledStudent: %v", len(marshalledStudent))
}
