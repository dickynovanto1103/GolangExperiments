package main

import (
	"context"
	"log"
	"time"
)

func runWithContext(ctx context.Context, sleepDur time.Duration) {
	ch := make(chan struct{})
	go func() {
		time.Sleep(sleepDur)
		ch <- struct{}{}
	}()
	log.Printf("hello now: %v", time.Now())
	log.Printf("ctx: %v", ctx)
	deadline, ok := ctx.Deadline()
	log.Printf("deadline: %v ok: %v", deadline, ok)
	select {
	case <-ch:
		log.Printf("done with the job")
	case <-ctx.Done():
		log.Printf("done with err: %v", ctx.Err())
	}
}

func main() {
	//server := &http.Server{
	//	Addr:              "localhost:8080",
	//	ReadTimeout:       time.Duration(1)*time.Second,
	//	WriteTimeout:      time.Duration(1)*time.Second,
	//}
	//
	//server.Shutdown()

	ctx1, cf := context.WithTimeout(context.Background(), 2*time.Second)
	log.Printf("CTX1")
	defer cf()
	log.Printf("done")
	<-ctx1.Done()
	log.Printf("abis done")

	ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelFunc()
	runWithContext(ctx, 10*time.Millisecond)
}