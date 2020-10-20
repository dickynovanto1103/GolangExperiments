package main

import (
	"fmt"
	"log"
)

type person struct {
	name *string
	age  *int
}

func NewPerson(name string, age int) *person {
	return &person{
		name: &name,
		age:  &age,
	}
}

//use %+v to print the field name --> it will print like this name:<val> age:<val>
func main() {
	person := NewPerson("dicky", 12)
	log.Printf("%+v", person)
	num := 123
	log.Printf("%+v", num) //there is no difference if we only want to print a native variable

	numString := fmt.Sprintf("%.8d", num)
	log.Printf("numString: %v", numString)
}
