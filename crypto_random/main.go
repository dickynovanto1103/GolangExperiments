package main

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

func main() {
	b := make([]byte, 32)
	n, err := rand.Read(b)
	log.Println(n, err, b) //it will create random elements array of byte
	s := base64.StdEncoding.EncodeToString(b)
	log.Printf("s: %v", s)
}

func createRandomNumber() {
	
}