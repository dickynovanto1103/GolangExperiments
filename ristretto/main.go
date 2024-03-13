package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"

	"github.com/dgraph-io/ristretto"
)

func main() {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters:        1e7,
		MaxCost:            1e6,
		BufferItems:        64,
		Metrics:            true,
		IgnoreInternalCost: true,
	})
	if err != nil {
		log.Panicf("error in creating cache, err: %v", err)
	}

	defer ants.Release()
	var wg sync.WaitGroup

	maxIdx := 100000
	for i := 0; i <= maxIdx; i++ {
		i := i
		wg.Add(1)
		ants.Submit(func() {
			defer wg.Done()
			key := fmt.Sprintf("asdfasdfadsfasdfasfasdfasdfadsfadsf%v", i)
			cache.SetWithTTL(key, fmt.Sprintf("vaadsfadsfadsfafdasdfasflue%v", i), 1, 3*time.Second)
		})
	}

	wg.Wait()

	log.Printf("done setting with ttl")

	//cache.Wait()
	log.Printf("cache metrics: %+v", cache.Metrics)

	var value interface{}
	var ok bool
	now := time.Now()
	for !ok {
		value, ok = cache.Get(fmt.Sprintf("asdfasdfadsfasdfasfasdfasdfadsfadsf%v", maxIdx))
	}

	log.Printf("%v %v took %v duration to set to cache", value, ok, time.Now().Sub(now))
}
