package main

import (
	"log"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	var x int = 42

	value := reflect.ValueOf(x)
	log.Printf("value: %+v", value)

	typ := reflect.TypeOf(x)
	log.Printf("type of x: %+v", typ)
	log.Printf("kind of value: %+v type of value: %v, string of value: %v", value.Kind(), value.Type(), value.String())

	var str string = "hello"
	valueStr := reflect.ValueOf(str)
	log.Printf("type: %v, value: %v, can complex: %v", valueStr.Type(), valueStr.String(), valueStr.CanAddr())
	log.Printf("is valid: %v, can set: %v", valueStr.IsValid(), valueStr.CanSet())

	elemAndSetStringExperiment()
}

// note:
// 1. I learnt that Elem must be called from a pointer or interface variable, if not, panic will occur
//   - And if it's pointer, then Elem will return the non pointer value of it.
func elemAndSetStringExperiment() {
	person := &Person{
		Name: "dicky",
		Age:  24,
	}

	val := reflect.ValueOf(person)
	val.UnsafePointer()
	elem := val.Elem()
	log.Printf("val: %v elem: %v", val, elem)

	nameValue := elem.FieldByName("Name")
	log.Printf("name value: %v", nameValue)
	if nameValue.IsValid() && nameValue.CanSet() {
		nameValue.SetString("angel")
	}

	log.Printf("person: %+v", person)
}
