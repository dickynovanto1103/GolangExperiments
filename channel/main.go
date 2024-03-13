package main

import "log"

func main() {
	ch := make(chan int)
	close(ch)

	ch <- 1

	for i := 0; i < 4; i++ {
		val := <-ch
		log.Printf("val: %d\n", val)
	}

}
