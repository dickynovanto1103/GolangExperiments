package main

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func main() {
	option := &redis.Options{
		Addr: ":6379",
	}
	rd := redis.NewClient(option)
	resSet := rd.HSet(context.Background(), "keyhash", "hello", "a123", "hello1", "233")
	log.Println("resSet:", resSet)
	res := rd.HGet(context.Background(), "keyhash", "hello")
	//res1 := rd.HGet(context.Background(), "keyhash", "hello1")
	resInt, err := res.Int64()
	log.Printf("res: %v %v", resInt, err)
}
