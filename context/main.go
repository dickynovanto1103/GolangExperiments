package main

import (
	"context"
	"log"
	"time"
)

func main() {
	//server := &http.Server{
	//	Addr:              "localhost:8080",
	//	ReadTimeout:       time.Duration(1)*time.Second,
	//	WriteTimeout:      time.Duration(1)*time.Second,
	//}
	//
	//server.Shutdown()
	ctx, cancelFunc := context.WithCancel(context.Background())

	go func(){
		log.Println("before sleep")
		time.Sleep(2*time.Second)
		cancelFunc()
		log.Printf("func done")
	}()

	<-ctx.Done()
	log.Printf("ctx done")
}