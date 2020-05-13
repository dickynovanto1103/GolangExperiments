package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func exp1() {
	ch := make(chan int)
	go func() {
		fmt.Println("hello world1")
		ch <- 1
	}()
	go func() {
		fmt.Println("hello world")
	}()
	go func() {
		fmt.Println("hello world2")
	}()
	//go func() {
	//	fmt.Println("hello world3")
	//}()

	val := <-ch
	fmt.Println("val:", val)
}

func exp2() {
	fmt.Println("main start")

	go func() {
		i := 0
		for {
			fmt.Println(i)
			i++
		}
	}()

	time.Sleep(2*time.Millisecond)

	fmt.Println("main ends")
	time.Sleep(1*time.Second)
}

func CanWeDoParallelOperation() {
	now := time.Now()
	threads := runtime.GOMAXPROCS(0)
	bil := int32(0)
	ch := make(chan int, threads)

	for i := 0; i < threads; i++ {
		go func(threadNum int) {
			//log.Println("processong thread num:", threadNum)
			for i := 0; i < 100000000; i++ {
				//atomic.AddInt32(&bil, 1)
				bil++
			}
			//log.Printf("thread num: %v done\n", threadNum)
			ch <- 1
		}(i)
	}

	for i:=0;i<threads;i++ {
		log.Println("done: ", i)
		<-ch
	}
	log.Println("totBil: ", bil)
	log.Println("time elapsed: ", time.Since(now))
}

func SerialAdd() {
	now := time.Now()
	var bil int32 = 0
	threads := runtime.GOMAXPROCS(0)
	for i:=0;i<threads;i++ {
		for j:=0;j<100000000;j++ {
			//atomic.AddInt32(&bil, 1)
			bil++
		}
	}
	log.Println("bil final: ", bil)
	log.Println("time elapsed: ", time.Since(now))
}
var n = 10000000

var arr1 = make([]int, n)
var arr2 = make([]int, n)
var sum = make([]int, n)

func SumParallel(n int) {
	now := time.Now()
	threads := runtime.GOMAXPROCS(0)
	ch := make(chan int, threads)
	//wg := sync.WaitGroup{}
	//wg.Add(threads)
	for i:=0;i<threads;i++ {
		go func(awal, akhir int) {
			for i:=awal;i<akhir;i++ {
				sum[i] = arr1[i] + arr2[i]
			}
			ch <- 1
			//wg.Done()
		}(i*n/threads, (i+1)*n/threads)
	}

	//wg.Wait()

	for i:=0;i<threads;i++ {
		<-ch
	}

	log.Println("parallel: ", time.Since(now))
}

func SumSerial(n int) {
	now := time.Now()

	for i:=0;i<n;i++{
		sum[i] = arr1[i] + arr2[i]
	}

	log.Println("serial: ", time.Since(now))
}

func gen(n int) {
	for i:=0;i<n;i++ {
		arr1[i] = i
		arr2[i] = i
	}

}

func main() {
	//exp1()
	//exp2()
	gen(n)
	log.Println("done")
	SumSerial(n)
	SumParallel(n)
	//CanWeDoParallelOperation()
	//SerialAdd()
}