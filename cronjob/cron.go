package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

func printEvery(time time.Duration) {
	fmt.Printf("executed every %v\n", time)
}

func main() {
	c := cron.New()
	c.AddFunc("@every 2s", func() {fmt.Println("executed every 2s")})
	c.Start()
	for {
		time.Sleep(1*time.Second)
	}
}
