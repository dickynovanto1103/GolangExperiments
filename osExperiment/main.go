package main

import (
	"log"
	"os"
)

func main() {
	uid := os.Getuid()
	log.Printf("uid: %v", uid)

	log.Printf("temp dir: %v", os.TempDir())
}
