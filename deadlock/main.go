package main

import (
	"log"
	"sync"
	"time"
)

type Something struct {
	val  int
	lock sync.Mutex
}

func main() {
	a := Something{
		val: 12,
	}

	b := Something{
		val: 13,
	}

	var wg sync.WaitGroup

	sum := func(a, b *Something) {
		defer wg.Done()

		a.lock.Lock()
		defer a.lock.Unlock()

		time.Sleep(1 * time.Second)

		b.lock.Lock()
		defer b.lock.Unlock()

		log.Printf("sum: %d\n", a.val+b.val)
	}

	wg.Add(2)

	go sum(&a, &b)
	go sum(&b, &a)

	wg.Wait()

	log.Printf("done")
}
