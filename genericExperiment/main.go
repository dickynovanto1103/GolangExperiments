package main

import (
	"log"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer
}

func SumNumber[A Number](a, b A) A {
	return a + b
}

func main() {
	log.Printf("%v", SumNumber(23, 23))
	log.Printf("%v", SumNumber(23.3, 23.3))
	log.Printf("%v", SumNumber(-9.0, 9.4))
}
