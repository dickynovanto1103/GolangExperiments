package main

import (
	"fmt"
	"go.uber.org/ratelimit"
	"time"
)

func main() {
	rl := ratelimit.New(10000) //per second
	prev := time.Now()
	for i:=0;i<100;i++{
		now := rl.Take()
		fmt.Println(i, now.Sub(prev))
		prev = now
	}

	//each operation will take 10 ms since the rate limit is 100 operations per second, it means that 1 operation will need 10 ms
}
