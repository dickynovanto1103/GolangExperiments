package main

import (
	"log"

	student2 "github.com/dickynovanto1103/GolangExperiments/protobufRedis/student"
	redis "github.com/go-redis/redis/v7"
	"github.com/gogo/protobuf/proto"
)

func main() {
	student := &student2.Student{
		Id:   proto.Int32(1),
		Name: proto.String("dicky"),
		IdCard: &student2.IDCard{
			Id: proto.Int32(123),
		},
		Books: []*student2.Book{
			{
				Id:   proto.Int32(12),
				Name: proto.String("clean architecture"),
			},
			{
				Id:   proto.Int32(11123),
				Name: proto.String("clean code"),
			},
		},
	}
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Panicf("error in ping, err: %v", err)
	}

	jsonString, errJsonMarshall := proto.Marshal(student)
	if errJsonMarshall != nil {
		log.Panicf("error in json marshall, err: %v", errJsonMarshall)
	}
	log.Printf("jsonString: %v", jsonString)

	nilai, err := client.Set("STUDENT_1", jsonString, -1).Result()
	if err != nil {
		log.Panicf("error in set key, err: %v", err)
	}
	log.Printf("nilai: %v", nilai)

	val, err := client.Get("STUDENT_1").Result()
	studentBaru := &student2.Student{}
	errUnmarshall := proto.Unmarshal([]byte(val), studentBaru)
	log.Printf("errorUnmarshall: %v", errUnmarshall)
	log.Printf("val: %+v, err: %v", studentBaru, err)
}
