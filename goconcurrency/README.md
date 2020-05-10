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