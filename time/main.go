package main

import (
	"log"
	"time"
)

func main() {
	date := time.Unix(10, 0)
	log.Println("unix time:", date.Unix())
}