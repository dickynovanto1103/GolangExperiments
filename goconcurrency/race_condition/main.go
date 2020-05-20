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

func main() {
	exp1()
}