package main

import (
	"hash/fnv"
	"log"
)

func main() {
	hashAlgo := fnv.New32a()
	key := "dicky"
	hashAlgo.Write([]byte(key))
	log.Printf("hash value: %v", hashAlgo.Sum32())

	value := 0x7fffffff
	log.Printf("value: %v", value)
}
