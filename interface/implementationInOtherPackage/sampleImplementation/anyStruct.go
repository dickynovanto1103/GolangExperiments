package sampleImplementation

import "fmt"

type AnyStruct struct {

}

func (s *AnyStruct) Execute() {
	fmt.Println("calling execute of any struct")
}