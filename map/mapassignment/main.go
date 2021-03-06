package main

import (
	"fmt"
	"log"
)

func change(m map[string]string) {
	m["dicky"] = "novanto"
}

func main() {
	mapper := make(map[string]string)
	change(mapper)
	fmt.Println("map: ", mapper)
	fmt.Println("mapper dari dicky: ", mapper["dicky"])

	otherMapper := copyMap(mapper)
	fmt.Println(otherMapper)
	mapper["dicky"] = "andre"

	fmt.Println("otherMapper:", otherMapper)
	fmt.Println("mapper:", mapper)

	resMap := returnMap()
	for k, v := range resMap {
		log.Println(k, v)
	}

	var cobaMap map[uint64]uint64
	log.Printf("cobaMap: %v", cobaMap)
	if cobaMap == nil {
		log.Printf("coba map nil")
	}
	if _, ok := cobaMap[1]; !ok {
		log.Printf("ga ketemu angka 1")
	}
}

func copyMap(mapper map[string]string) map[string]string {
	otherMap := make(map[string]string)

	for k, v := range mapper {
		otherMap[k] = v
	}

	return otherMap
}

func returnMap() (res map[int]string) {
	res = make(map[int]string)
	res[123] = "dicky"
	return res
}
