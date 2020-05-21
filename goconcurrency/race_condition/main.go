package main

import (
	"fmt"
)

func exp1() {
	done := make(chan bool)
	mapper := make(map[int]int)
	//mutex := sync.Mutex{}

	go func() {
		//mutex.Lock()
		mapper[1] = 1
		//mutex.Unlock()
		done <- true
	}()
	//mutex.Lock()
	mapper[2] = 2
	//mutex.Unlock()
	<-done
	for k, v := range mapper {
		fmt.Printf("k: %v v: %v\n", k, v)
	}
}

func raceOnLoopCounter() {
	ch := make(chan int)
	for i:=0;i<5;i++ {
		go func(j int) {
			fmt.Printf("%v", j)
			ch <- 1
		}(i)
	}
	for i:=0;i<5;i++ {
		<-ch
	}
}

func main() {
	//exp1()
	raceOnLoopCounter()
}