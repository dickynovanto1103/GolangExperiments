package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/dickynovanto1103/GolangExperiments/rpc/server/arith"
)

func main() {
	arith := new(arith.Arith)
	err := rpc.Register(arith)
	if err != nil {
		log.Fatalf("error in registering type arith, err: %v", err)
	}
	rpc.HandleHTTP()
	listener, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatalf("error in getting listener, err: %v", err)
	}
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatalf("error serving http")
	}
}
