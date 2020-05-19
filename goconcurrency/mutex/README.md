# Mutex experiment note
1. We can just put sync.Mutex inside the struct to inherit Lock() and Unlock() method
2. Tips in using mutex:
    1. Make sure we only lock a certain critical path, don't lock part of code that requires I/O operation
    2. We can defer unlock if in the code, there are many return. This it to prevent unlocked lock + make code more tidy
    3. We cannot use defer in for loop, will cause deadlock
3. If we lock 2 times, then it will cause deadlock --> locking 2 times
4. If we lock 1 times and no unlocking happens, the deadlock will not happen
5. We can use RWMutex if the case is that we want to do read and write operation.
    1. Why can't we use the usual Mutex for this case? because in order to get the correct value when we read, we need to do locking first. In some cases, the read operation can require more than 1 times locking. Therefore we can use RLock() 
    