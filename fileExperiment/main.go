package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	experimentWriteToFileUsingFPrintf()
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
