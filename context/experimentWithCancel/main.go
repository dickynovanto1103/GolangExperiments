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
		resultCh <- idx
		idx++
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	ctx := context.Background()
	ctx, cancelFunc := context.WithCancel(ctx)
	log.Printf("cancel context is created")

	resultCh := make(chan int, 10)


	go handleResults(ctx, resultCh)
	go process(ctx, resultCh)

	log.Printf("concurrently running process and handleResults")
	time.Sleep(500 * time.Millisecond)
	log.Printf("cancelFunc is called")
	cancelFunc()

	time.Sleep(1 * time.Second)

	log.Printf("done")
}