package main

import (
	"fmt"
	"log"
)

func main() {
	defer func() {
		fmt.Println("halo dunia")
	}()

	log.Fatal("error fatal here") //if we use this, the defer function will not be executed after log.Fatal is executed
	// log.Panic("error panic here") //instead we should use other log methods, like log.Panic
}