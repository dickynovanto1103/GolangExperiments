## Database Transaction
1. We can use TCC (Try confirm cancel)
2. Example of TCC usage:
    - Try to deduct a money
    - Then confirm / cancel if there is an error
    - If confirm / cancel fails, we have TCC checker that get list of pending transactions, then try to retry the transaction
3. There is a go-saga pattern also

## Saga Pattern
1. Divide a transaction into several sub transactions (T1, T2, ..., Tn)
    - Example: for booking a trip, there is hotel booking and car booking, so we divide the trip booking transaction into car booking and hotel booking sub-transactions
2. Each sub transaction should have compensation logic (in case of the sub-transaction fails). The compensation should be idempotent
3. So it will be either:
    1. `T1, T2, ..., Tn` or  
    2. `T1, T2, ..., Ti, Ci, Ci-1, ... , C1`: backward recovery method
4. This saga pattern will ensure the system will still be in consistent state. But we must sacrifice atomicity for high availability
5. There is SEC (Saga Execution Coordinator), this coordinator will control the sub-transactions and the compensations
6. What is compensations fails? just retry until it is fine since the compensation is idempotent 