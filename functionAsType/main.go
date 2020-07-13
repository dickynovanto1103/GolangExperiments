package main

import (
	"fmt"
	"log"
)

type Func func(s string) error

type convert func(int) string

func value(x int) string {
	return fmt.Sprintf("%v", x)
}

func coba(p Func, s string) {
	p(s)
}

func main() {
	var s string
	bil := 123
	s = value(bil)
	s += "12"

	log.Printf("s: %v", s)
	coba(func(st string) error {
		fmt.Printf(st)
		return nil
	}, s)
}