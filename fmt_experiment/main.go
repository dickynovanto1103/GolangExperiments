package main

import (
	"fmt"
	"log"
)

func main() {
	//ScanlnExperimentInputName()
	//ScanLnInputInteger()
	//ScanLnInput2Integers()
	//ScanInputInteger()
	//ScanInput2Integers()
	trySscanOfName()
}

// Scanln will read input until \n
// The n there is the number of variable being processed. If input is 1 var, it returns 1, if 2, then n = 2
func ScanlnExperimentInputName() {
	var name string
	fmt.Printf("enter your name!\n")
	n, err := fmt.Scanln(&name)
	if err != nil {
		log.Panicf("error in scanning name, err: %v", err)
	}
	log.Printf("n: %v", n)

	log.Printf("hello %s", name)
}

func ScanLnInputInteger() {
	var num int
	fmt.Printf("enter a number!\n")
	//if we input a string, it will return error: expexted integer
	n, err := fmt.Scanln(&num)
	if err != nil {
		log.Panicf("error in inputing number, err: %v", err)
	}
	log.Printf("n: %v", n)

	log.Printf("%v * 2 = %v", num, num*2)
}

// for this, we must input the 2 numbers separated with space, we cannot input like this:
// 12 (new line) 13 (new line)...it will return err: unexpected new line
func ScanLnInput2Integers() {
	var a, b int
	fmt.Printf("enter 2 numbers!\n")
	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		log.Panicf("error in inputing number, err: %v", err)
	}
	log.Printf("n: %v", n)

	log.Printf("%v", a+b)
}

func ScanInputInteger() {
	var num int
	n, err := fmt.Scan(&num)
	if err != nil {
		log.Panicf("error in fmt.Scan integer, err: %v", err)
	}
	fmt.Printf("n: %v", n)
	fmt.Printf("num inputted: %v", num)
}

// oh for scan, if we input 2 integers...we can do it these ways:
// 12 13, or
// 12 (new line) 13 (new line)
// the new line will be treated the same as space
func ScanInput2Integers() {
	var a, b int
	n, err := fmt.Scan(&a, &b)
	if err != nil {
		log.Panicf("error in fmt.Scan integer, err: %v", err)
	}
	fmt.Printf("n: %v\n", n)
	fmt.Printf("sum: %v", a+b)
}

func trySscanOfName() {
	var firstName, lastName string
	str := "john doe"
	n, err := fmt.Sscan(str, &firstName, &lastName)
	if err != nil {
		log.Panicf("error in scanning, err: %v", err)
	}

	fmt.Printf("n: %v\n", n)
	fmt.Printf("firstname: %v lastName: %v", firstName, lastName)
}
