package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("hello world1")
		ch <- 1
	}()
	go func() {
		fmt.Println("hello world")
	}()

	val := <-ch
	fmt.Println("val:", val)
}