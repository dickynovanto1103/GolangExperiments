package main

import (
	"log"
	"sync"
)

func runPanic(i int) {
	defer func() {
		log.Printf("in defer")
	}()
	log.Panicf("dicky nakal, idx: %v", i)
}

func main() {
	defer func() {
		log.Printf("done after panic?")
	}()
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer func ()  {
				log.Printf("defer func inside a goroutine idx: %v", i)	
			}()
			defer wg.Done()
			runPanic(i)
		}(i)
	}

	wg.Wait()

	log.Printf("done")
}
