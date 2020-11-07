package main

import (
	"context"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
)

func experimentSentryRecover() {
	//just take a look at the sentry.Recover() implementation and we will be able to understand

	//defer func() {
	//	if err := recover(); err != nil {
	//		sentry.CurrentHub().Recover(err)
	//	} else {
	//		log.Printf("received nil err")
	//		sentry.CurrentHub().Recover(err)
	//	}
	//	sentry.Recover()
	//
	//	log.Printf("done recover")
	//}()
	ctx := context.WithValue(context.Background(), 2, "test ctx value")
	defer sentry.RecoverWithContext(ctx)

	log.Printf("test panic")
	cobaPanic("halo test panic3")
}

func cobaPanic(msg string) {
	panic(msg)
}

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              "http://1e4058a3ee6f403ab796a8e17f0e355f@localhost:8080/1",
		AttachStacktrace: true,
	})

	if err != nil {
		log.Fatalf("error sentry init, err: %v", err)
	}
	log.Printf("no error")
	experimentSentryRecover()

	//eventId := sentry.CaptureException(errors.New("hello error"))
	//log.Printf("eventId: %v", *eventId)
	defer sentry.Flush(5 * time.Second)
	//log.Printf("done flushing")
}
