package main

import (
	"fmt"
	"log"
)

func exp1(){
	ch := make(chan bool)
	done := make(chan bool)
	go func() {
		select{
		case val := <-ch:
			fmt.Println("val ch: ", val)
			done <- true
		}
	}()

	close(ch)
	ch = make(chan bool)
	close(ch)
	<-done
}

func exp2() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for{
			job, more := <-jobs
			if !more {
				log.Println("finished processing all jobs")
				done <- true
				return
			}
			log.Printf("received job: %v", job)
		}
	}()

	for i:=0;i<10;i++ {
		jobs <- i
		log.Println("sent jobs:", i)
	}
	close(jobs)
	log.Println("sent all jobs")
	<-done
}

func main() {
	//exp1()
	exp2()
}
