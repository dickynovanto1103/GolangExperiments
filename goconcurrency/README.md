# Go Concurrency Note
1. If we have this piece of code

```
    ch := make(chan int)
    go func() {
        fmt.Println("hello world1")
        ch <- 1
    }()
    go func() {
        fmt.Println("hello world")
    }()
    
    val := <-ch
    fmt.Println("val:", val)
```
Since it request the value of the ch, then it will start to seek in all goroutines, which one that has sending to ch channel operation.
The order of it is going from the latest goroutine created, till the first. So it's like stack. If all goroutines are observed and no send message to ch channel operation, then all goroutines are asleep and it will cause deadlock.

2. From exp2() function, we can learn that once the main goroutines sleep time is over, it will continue to do the main flow, so the infinite loops in the function that the other goroutine is working on will stop.
3. Closing channel is beneficial for tell that this channel will not be sending anymore / finished processing all request, so on the receiver side, need to handle the case when the channel has been closed.
4. From Rob Pike's presentation, I learnt that concurrency != parallelism, concurrency enables parallelism.
    1. We can design the algorithm such that it can be executed as an independent piece of algorithm, so that it can be further parallelized
    2. In real life, we can build a load balancer that dispatch a lot of requests to different workers and it can be done in parallel fashion. 