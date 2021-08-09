package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	uid := os.Getuid()
	log.Printf("uid: %v", uid)

	log.Printf("temp dir: %v", os.TempDir())

	file, err := ioutil.TempFile("./", "dicky_*")
	if err != nil {
		log.Panicf("error in creating temporary file, err: %v", err)
	}
	defer file.Close()
	log.Printf("filename: %v", file.Name())
	oldPath := fmt.Sprintf("%v", file.Name())
	os.Rename(oldPath, "dicky.txt")
	pid := os.Getpid()
	log.Printf("pid: %v", pid)

}
