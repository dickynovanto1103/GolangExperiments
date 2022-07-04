package main

import (
	"context"
	"fmt"
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

func testContext(ctx context.Context, cancelFunc context.CancelFunc) {
	done := make(chan struct{})
	for i:=0;i<5;i++{
		//log.Printf("i: %v", i)
		go func(ctx context.Context, i int) {
			//log.Printf("run forever loop")
			for {
				select {
				case <-ctx.Done():
					done <- struct{}{}
					log.Printf("test context %v done, err: %v",i, ctx.Err())
					return
				default:
					fmt.Println("for loop")
				}
				time.Sleep(1000*time.Millisecond)
				//log.Printf("1000 ms\n")
			}
		}(ctx, i)
	}

	go func() {
		//log.Printf("sleep for 3 seconds")
		time.Sleep(3*time.Second)
		//log.Printf("cancel func")
		cancelFunc()
	}()
	//NOTE: if we want to see the
	for i:=0;i<5;i++{
		<-done
	}
	time.Sleep(1*time.Second)
	log.Printf("done here")
}

func doSomething(ctx context.Context) {
	log.Printf("do something, ctx value: %v, andre: %v", ctx.Value("dicky"), ctx.Value("andre"))
}

func experimentWithValue() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "dicky", "ganteng")
	ctx = context.WithValue(ctx, "dicky", "sekali")
	ctx = context.WithValue(ctx, "andre", "ganteng jg")
	context.WithCancel()
	ctx.Err()
	doSomething(ctx)
}

func main() {
	experimentWithValue()
	//server := &http.Server{
	//	Addr:              "localhost:8080",
	//	ReadTimeout:       time.Duration(1)*time.Second,
	//	WriteTimeout:      time.Duration(1)*time.Second,
	//}
	//
	//server.Shutdown()

	//ctx1, cf := context.WithTimeout(context.Background(), 2*time.Second)
	//log.Printf("CTX1")
	//defer cf()
	//log.Printf("done")
	//<-ctx1.Done()
	//log.Printf("abis done")
	//
	//ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	//defer cancelFunc()
	//runWithContext(ctx, 3*time.Second)

	// ctx2, cancelFunc2 := context.WithCancel(context.Background())
	// defer cancelFunc2()
	
	//log.Printf("\n\n\n")
	//log.Printf("start of test context")
	
	// testContext(ctx2, cancelFunc2)
}