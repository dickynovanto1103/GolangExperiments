package main

import "fmt"

func main() {
	ch := make(chan bool)
	done := make(chan bool)
	go func() {
		select{
		case val := <-ch:
			fmt.Println("val ch: ", val)
			done <- true
		}
	}()

	fmt.Println("halting")
	close(ch)
	ch = make(chan bool)
	close(ch)
	<-done
}
