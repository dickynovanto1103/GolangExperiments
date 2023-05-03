package main

import (
	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(1, 3)

func main() {
	//newHandler := NewHandler()
	limiter.Wait()
}
