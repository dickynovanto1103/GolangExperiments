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
	nonPointerObjThenTryToSetString()
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
	elem := val.Elem()
	log.Printf("val: %v elem: %v", val, elem)

	nameValue := elem.FieldByName("Name")
	if nameValue.IsValid() && nameValue.CanSet() {
		nameValue.SetString("angel")
	}

	log.Printf("person: %+v", person)
}

// if object is a non pointer object, then the field.CanSet() will return false, and if we
// try to do field.SetString / SetInt / etc, it will panic
func nonPointerObjThenTryToSetString() {
	person := Person{
		Name: "DICKY",
		Age:  24,
	}

	value := reflect.ValueOf(person)
	log.Printf("value: %+v", value)
	nameField := value.FieldByName("Name")
	log.Printf("can set name field: %v", nameField.CanSet())
	//nameField.SetString("asu") //it will panic
	person.Name = "dickyyy"
	log.Printf("person: %+v", person)
	val := reflect.ValueOf(person)
	log.Printf("val: %+v", val)
}
