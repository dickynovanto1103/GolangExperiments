package main

import (
	"runtime"
	"sync"
)

type st struct {
	sync.Mutex
	bil int
}

var str = &st{}

func increment(times int) {
	for i:=0;i<times;i++ {
		str.Lock()
		str.bil++
		str.Unlock()
	}
}

func Pattern1Mutex() {
	threads := runtime.GOMAXPROCS(0)
	ch := make(chan int)

	for i:=0;i<threads;i++ {
		go func() {
			increment(10000)
			<-ch
		}()
	}

	for i:=0;i<threads;i++ {
		ch <- 1
	}
}

func Pattern2Mutex() {
	for i:=0;i<2;i++ {
		str.Lock()
	}
}

func main() {
	//Pattern1Mutex()
	//log.Printf("final bil: %v", str.bil)

	Pattern2Mutex()
}