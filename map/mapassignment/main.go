package main

import "fmt"

func change(m map[string]string) {
	m["dicky"] = "novanto"
}

func main() {
	mapper := make(map[string]string)
	change(mapper)
	fmt.Println("map: ", mapper)
	fmt.Println("mapper dari dicky: ",mapper["dicky"])

	otherMapper := copyMap(mapper)
	fmt.Println(otherMapper)
	mapper["dicky"] = "andre"

	fmt.Println("otherMapper:", otherMapper)
	fmt.Println("mapper:", mapper)
}

func copyMap(mapper map[string]string) map[string]string {
	otherMap := make(map[string]string)

	for k, v := range mapper {
		otherMap[k] = v
	}

	return otherMap
}