package main

import (
	"log"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan int)
	go runTickerUntilDone(ticker, done)

	log.Printf("go routine created")
	time.Sleep(3 * time.Second)
	done <- 1
}

func runTickerUntilDone(ticker *time.Ticker, done chan int) {
	for {
		select {
		case <-done:
			log.Printf("done")
			return
		case t := <-ticker.C:
			log.Printf("tick at t: %v, in seconds: %v", t, time.Now().Unix())
		}
	}
}
