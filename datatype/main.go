package main

import (
	"log"
	"math"
)

func main() {
	var a int
	var ab int64 = math.MaxInt64
	a = int(ab) //in 64-bit machine, this can be done as int here is allocated 64 bits

	log.Printf("a: %v ab: %v", a, ab)
}
