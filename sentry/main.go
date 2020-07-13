package main

import (
	"errors"
	"github.com/getsentry/sentry-go"
	"log"
	"time"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "http://668b1a9326af47fbb032e251a5c44550@localhost:8080/2",
	})

	if err != nil {
		log.Fatalf("error sentry init, err: %v", err)
	}
	log.Printf("no error")

	eventId := sentry.CaptureException(errors.New("hello error"))
	log.Printf("evenId: %v", *eventId)
	sentry.Flush(5*time.Second)
}
