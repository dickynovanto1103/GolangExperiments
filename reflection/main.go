package main

import (
	"log"
	"reflect"
)

func getFullName(i interface{}) string {
	value := reflect.ValueOf(i)
	if value.Kind() != reflect.Func {
		return ""
	}

	return "halo"
}

func main() {
	ans := getFullName(".PriceCheck")
	log.Println("ans: ", ans)
}