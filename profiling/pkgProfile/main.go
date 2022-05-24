package main

import (
	"io"
	"log"
	"os"
)

func readByte(r io.Reader) (rune, error) {
	var bytes [1]byte
	_, err := r.Read(bytes[:])
	return rune(bytes[0]), err
}

func main() {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Panicf("error in opening file: %v, err: %v", fileName, err)
	}
	// bytes, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	log.Panicf("error in reading all from file, err: %v", err)
	// }
	// log.Printf("bytes: %v", string(bytes))
	for i := 0; i < 100; i++ {
		c, err := readByte(file)

		log.Printf("c: %v err: %v", string(c), err)
	}
	file.Read()

}
