# Race Condition note
1. When we do exp1, and we run: `go run -race main.go`, we will get warning on race condition because we do concurrent write operation to the non-thread safe map
    1. In order to prevent that, we can use mutex.Lock()
2. We can print the data race warning into a file by setting env var `GORACE="log_path=<filepath>"`
