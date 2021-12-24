package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	// callAPIWithoutTimeout()
	callAPIWithTimeout()
}

func callAPIWithTimeout() {
	httpClient := &http.Client{
		Timeout: 1 * time.Second,
	}
	resp, err := httpClient.Get("http://localhost:8080/slow")
	logRespAndErr(resp, err)
}

// func callAPIWithoutTimeout() {
// 	resp, err := http.Get("http://localhost:8080/slow")
// 	logRespAndErr(resp, err)
// }

func logRespAndErr(resp interface{}, err error) {
	log.Printf("resp: %+v err: %v", resp, err)
}
