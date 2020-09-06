package main

import (
	"log"
	"net/rpc"

	"github.com/dickynovanto1103/GolangExperiments/rpc/server/arith"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatalf("error dial HTTP from client, err: %v", err)
	}
	args := &arith.Args{
		A: 100,
		B: 5,
	}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatalf("error client call, err: %v", err)
	}
	log.Printf("reply: %v", reply)

	var replyDiv arith.Quotient
	divCall := client.Go("Arith.Divide", args, &replyDiv, nil)
	result := <-divCall.Done

	log.Printf("result: %+v, args: %+v, reply: %+v", result, result.Args, result.Reply.(*arith.Quotient))
}
