package main

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/panjf2000/ants/v2"
	"github.com/pkg/errors"
)

var sum int32

func myFunc(i int32) {
	atomic.AddInt32(&sum, i)
	runtime.Gosched()
	fmt.Printf("run with %d\n", i)
}

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

func main() {
	defer ants.Release()
	runTimes := 1000

	var wg sync.WaitGroup
	syncCalcSum := func() {
		demoFunc()
		wg.Done()
	}

	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		err := ants.Submit(syncCalcSum)
		if err != nil {
			log.Panicf("error in submitting task, err: %v", err)
		}
	}

	wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Printf("finish running all tasks")

	poolWithFunc, err := ants.NewPoolWithFunc(10, func(i interface{}) {
		myFunc(i.(int32))
		wg.Done()
	})
	if err != nil {
		log.Panicf("error in creating new pool with func, err: %v", err)
	}

	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		if err = poolWithFunc.Invoke(int32(i)); err != nil {
			log.Panicf("error in invoking, err: %v", err)
		}
	}
	wg.Wait()

	fmt.Printf("running goroutines: %d\n", poolWithFunc.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum)
	if sum != 499500 {
		panic("the final result is wrong!!!")
	}

	pool, err := ants.NewPool(10)
	if err != nil {
		log.Panicf("error in creating new pool, err: %v", err)
	}

	errChan := make(chan error)

	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		if err = pool.Submit(func() {
			err = returnErrorOrNot()
			errChan <- err
			wg.Done()
		}); err != nil {
			log.Panicf("error in submitting task, err: %v", err)
		}

	}

	wg.Wait()

}

func returnErrorOrNot() error {
	if rand.Intn(10) < 5 {
		return errors.New("error")
	}
	return nil
}
