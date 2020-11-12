package main

import (
	"log"
	"strings"
)

func experimentEqualFold() {
	a := "dicky Novanto"
	b := "DICKY NOVANTO"
	log.Printf("%v", strings.EqualFold(a, b))
}

func main() {
	experimentEqualFold()
}
