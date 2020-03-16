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
}
