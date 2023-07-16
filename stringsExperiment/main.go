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
	//experimentEqualFold()
	experimentStringsCount()
}

func experimentStringsCount() {
	a := "aaaaaa"
	log.Printf("%v", strings.Count(a, "aa"))
	strings.Index()
	strings.Fields()
	strings.IndexAny()
}
