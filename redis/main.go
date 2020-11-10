package main

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func testGetAndSetInteger(redisClient *redis.Client) {
	_, err := redisClient.Set(context.Background(), "dicky", 10, 10*time.Minute).Result()
	if err != nil {
		log.Fatalf("error set in cache, err: %v", err)
	}
	num, err := redisClient.Get(context.Background(), "dicky").Result()
	log.Printf("num: %v type: %T, err: %v", num, num, err)
}

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
	// redis transactions
	pipe := rd.TxPipeline()
	incr := pipe.Incr(context.Background(), "counter")
	pipe.Expire(context.Background(), "counter", 5*time.Second) //it will expire after 5 seconds this counter var is not used anymore
	_, err = pipe.Exec(context.Background())
	//time.Sleep(2 * time.Second)
	log.Printf("err: %v incr val: %v", err, incr.Val())
	testGetAndSetInteger(rd)
}
