package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"
	"github.com/bcicen/grmon/agent"
)

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	ch := make(chan int, 10)
	go func() {
		obj := make(map[string]float64)
		if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {
			log.Printf("err in decoding, err: %v", err)
			ch <- http.StatusBadRequest
			return
		}

		//simulate long processing
		time.Sleep(time.Duration(rand.Intn(400)) * time.Millisecond)
		ch <- http.StatusOK
	}()

	select {
	case status := <-ch:
		w.WriteHeader(status)
	case <-time.After(200 * time.Millisecond):
		w.WriteHeader(http.StatusRequestTimeout)
	}
}

func main() {
	grmon.Start()
	http.HandleFunc("/log", HandleFunc)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Panicf("error in listening and serving, err: %v", err)
	}
}
