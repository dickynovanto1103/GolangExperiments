package main

import (
	"context"
	"log"
	"time"
)

func handleResults(ctx context.Context, resultCh chan int) {
	for {
		select {
		case res := <-resultCh:
			log.Printf("res: %v", res)
		case <-ctx.Done():
			err := ctx.Err()
			if err == nil {
				log.Panicf("error should not be nil")
			}
			log.Printf("err is: %v", err)
			return
		}
	}
}

func process(ctx context.Context, resultCh chan int) {
	idx := 0;
	for{
		if ctx.Err() != nil {
			break
		}
		resultCh <- idx
		idx++
		time.Sleep(100 * time.Millisecond)
		log.Printf("still running")
	}
}

func main() {
	ctx := context.Background()
	ctx, cancelFunc := context.WithTimeout(ctx, 2 * time.Second)
	defer cancelFunc() //this is to release resources if operation is finished before timeout
	log.Printf("timeout context is created")

	resultCh := make(chan int, 10)


	go handleResults(ctx, resultCh)
	go process(ctx, resultCh)

	log.Printf("concurrently running process and handleResults")
	time.Sleep(500 * time.Millisecond)

	time.Sleep(3 * time.Second)

	log.Printf("done")
}