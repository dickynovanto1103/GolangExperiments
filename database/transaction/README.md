## Database Transaction
1. We can use TCC (Try confirm cancel)
2. Example of TCC usage:
    - Try to deduct a money
    - Then confirm / cancel if there is an error
    - If confirm / cancel fails, we have TCC checker that get list of pending transactions, then try to retry the transaction
3. There is a go-saga pattern also