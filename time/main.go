package main

import (
	"fmt"
	"log"
	"time"
)

func testSleep() {

	fmt.Println("halo")
	now := time.Now()
	time.Sleep(1) //equal to 1 microsecond
	then := time.Now()

	fmt.Println("dunia selisih", then.Sub(now))
}

func main() {
	date := time.Unix(10, 0)
	log.Println("unix time:", date.Unix())
	testSleep()
}