package main

import (
	"log"

	student2 "github.com/dickynovanto1103/GolangExperiments/protobufRedis/student"
	"github.com/go-redis/redis/v7"
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

	//experiment MGet
	keys := []string{"STUDENT_1", "STUDENT_2"}
	results, err := client.MGet(keys...).Result()
	if err != nil {
		log.Panicf("error in MGet, err: %v", err)
	}

	for _, res := range results {
		if res == nil {
			continue
		}
		resString := res.(string)

		val := []byte(resString)
		student := &student2.Student{}
		errUnmarshall = proto.Unmarshal(val, student)
		log.Printf("errUnmarshall: %v", errUnmarshall)
		log.Printf("student: %+v\n", student)
	}

	_, err = client.HSet("test", "STUDENT_1", jsonString, "STUDENT_2", jsonString).Result()
	if err != nil {
		log.Fatalf("fail to hset, err: %v", err)
	}
	mapRes, err := client.HGetAll("test").Result()
	for key, valString := range mapRes {
		val := []byte(valString)
		student := &student2.Student{}
		errUnmarshall = proto.Unmarshal(val, student)
		log.Printf("errUnmarshall: %v", errUnmarshall)
		log.Printf("test key: %v student: %+v\n", key, student)
	}
	anotherStudent := &student2.Student{
		Id:     proto.Int32(12),
		Name:   nil,
		IdCard: nil,
		Books:  nil,
	}
	marshalledStudent, _ := proto.Marshal(anotherStudent)
	_, err = client.HSet("test", "STUDENT_2", marshalledStudent).Result()
	mapRes, err = client.HGetAll("test").Result()
	for key, valString := range mapRes {
		val := []byte(valString)
		student := &student2.Student{}
		errUnmarshall = proto.Unmarshal(val, student)
		log.Printf("errUnmarshall: %v", errUnmarshall)
		log.Printf("again test key: %v student: %+v\n", key, student)
	}
	_, err = client.HSet("test1", "STUDENT", marshalledStudent).Result()
	pipe := client.Pipeline()
	pipe.HGetAll("test")
	res2, err := pipe.HGetAll("test1").Result()
	answer, err := pipe.Exec()
	if err != nil {
		log.Fatalf("error in exec, err: %v", err)
	}
	for _, ans := range answer {
		res := ans.(*redis.StringStringMapCmd)
		log.Printf("res: %+v", res)
		mapRes, err = res.Result()
		for key, jawab := range mapRes {
			log.Printf("key: %v, jawab: %+v", key, jawab)
			val := []byte(jawab)
			student := &student2.Student{}
			errUnmarshall = proto.Unmarshal(val, student)
			log.Printf("errUnmarshall: %v", errUnmarshall)
			log.Printf("again again test key: %v student: %+v\n", key, student)
		}
	}
	log.Printf("res1: %+v, res2: %+v answer: %+v", 12, res2, answer)
}
