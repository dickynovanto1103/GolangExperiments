package main

import (
	"log"
	"time"
)

func main() {

	for range time.Tick(0) {
		log.Printf("hello, time now: %v", time.Now().Unix())
	}
}
