package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go startWorker(i, &wg)
	}

	log.Printf("now waiting")
	wg.Wait()

	log.Printf("after wait")
}

func startWorker(workerNum int, wg *sync.WaitGroup) {
	log.Printf("worker: %v starts", workerNum)

	time.Sleep(1 * time.Second)

	log.Printf("worker: %v finished processing", workerNum)
	wg.Done()
}
