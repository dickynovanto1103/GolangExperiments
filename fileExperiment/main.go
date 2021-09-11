package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// experimentWriteToFileUsingFPrintf()
	experimentReadFromCertainOffsetInFile()
}

func experimentReadFromCertainOffsetInFile() {
	file, err := os.Open("./dicky")
	if err != nil {
		log.Panicf("error in opening file, err: %v", err)
	}

	offset := 1
	var bytes []byte = make([]byte, 100)
	n, err := file.ReadAt(bytes, int64(offset))
	if err != nil && err != io.EOF {
		log.Panicf("error in read at, err: %v", err)
	}
	log.Printf("n: %v bytes: %v, bytes in string: %v", n, bytes, string(bytes))
}

func experimentWriteToFileUsingFPrintf() {
	fileName := "dicky"
	file, err := os.Create(fileName)
	if err != nil {
		log.Panicf("error in creating file, err: %v", err)
	}
	for i := 0; i < 10; i++ {
		fmt.Fprintf(file, "asu %v\n", i)
	}

}
