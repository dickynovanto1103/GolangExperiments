package main

import (
	"io/ioutil"
	"log"
)

func main() {
	content := []byte("content here")
	tmpFile, err := ioutil.TempFile("./", "contoh")
	if err != nil {
		log.Fatalf("error in creating tempfile, err: %v", err)
	}
	if _, err := tmpFile.Write(content); err != nil {
		log.Printf("err in writing to tmpFile: %v", err)
		return
	}
	log.Printf("name: %v", tmpFile.Name())
	if err := tmpFile.Close(); err != nil {
		log.Printf("error in closing file, err: %v", err)
		return
	}
}
