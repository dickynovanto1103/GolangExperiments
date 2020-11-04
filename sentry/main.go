package main

import (
	"errors"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "http://1e4058a3ee6f403ab796a8e17f0e355f@localhost:8080/1",
	})

	if err != nil {
		log.Fatalf("error sentry init, err: %v", err)
	}
	log.Printf("no error")

	for i := 0; i < 1000000; i++ {
		eventId := sentry.CaptureException(errors.New("hello error"))
		log.Printf("eventId: %v", *eventId)
	}
	sentry.Flush(5 * time.Second)
}
