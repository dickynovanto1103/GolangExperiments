# Race Condition note
1. When we do exp1, and we run: `go run -race main.go`, we will get warning on race condition because we do concurrent write operation to the non-thread safe map
    1. In order to prevent that, we can use mutex.Lock()
2. We can print the data race warning into a file by setting env var `GORACE="log_path=<filepath>"`
## Typical Race Condition
1. Race on loop counter, as we can see in the function, we must pass the i to the parameter of the function in order to prevent the race on loop counter.
    - This happens because we have looped 5 times and the value of i now is 5, then when it starts printing i, it shows 5 for all 5 functions run by the goroutines