package main

import (
	"context"
	"github.com/lysu/go-saga"
	"log"
	"os"
	"time"
)

func main() {
	// 1. Define sub-transaction method, anonymous method is NOT required, Just define them as normal way.
	DeduceAccount := func(ctx context.Context, account string, amount int) error {
		// Do deduce amount from account, like: account.money - amount
		return nil
	}
	CompensateDeduce := func(ctx context.Context, account string, amount int) error {
		// Compensate deduce amount to account, like: account.money + amount
		return nil
	}
	DepositAccount := func(ctx context.Context, account string, amount int) error {
		// Do deposit amount to account, like: account.money + amount
		return nil
	}
	CompensateDeposit := func(ctx context.Context, account string, amount int) error {
		// Compensate deposit amount from account, like: account.money - amount
		return nil
	}

	saga.StorageConfig.Kafka.ZkAddrs = []string{"0.0.0.0:2181"}
	saga.StorageConfig.Kafka.BrokerAddrs = []string{"0.0.0.0:9092"}
	saga.StorageConfig.Kafka.Partitions = 1
	saga.StorageConfig.Kafka.Replicas = 1
	saga.StorageConfig.Kafka.ReturnDuration = 50 * time.Millisecond

	saga.AddSubTxDef("deduce", DeduceAccount, CompensateDeduce).
		AddSubTxDef("deposit", DepositAccount, CompensateDeposit)
	sagaId := uint64(2)
	from, to := "foo", "bar"
	amount := 100

	logger := log.New(os.Stdout, "saga_test", 1)
	saga.SetLogger(logger)

	saga.StartSaga(context.Background(), sagaId).
		ExecSub("deduce", from, amount).
		ExecSub("deposit", to, amount).
		EndSaga()
}