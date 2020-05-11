package main

import (
	"fmt"
	"time"
)

func exp1() {
	ch := make(chan int)
	go func() {
		fmt.Println("hello world1")
		ch <- 1
	}()
	go func() {
		fmt.Println("hello world")
	}()
	go func() {
		fmt.Println("hello world2")
	}()
	//go func() {
	//	fmt.Println("hello world3")
	//}()

	val := <-ch
	fmt.Println("val:", val)
}

func exp2() {
	fmt.Println("main start")

	go func() {
		i := 0
		for {
			fmt.Println(i)
			i++
		}
	}()

	time.Sleep(2*time.Millisecond)

	fmt.Println("main ends")
	time.Sleep(1*time.Second)
}

func main() {
	//exp1()
	exp2()
}