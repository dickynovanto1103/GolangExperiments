package main

import (
	"log"
	"sort"
)

type Person struct {
	Name *string
}

func main() {
	//dicky := "dicky"
	novanto := "novanto"
	persons := []*Person{
		{
			Name: &novanto,
		},
		{
			Name: nil,
		},
	}

	sort.Slice(persons, func(i, j int) bool {
		log.Printf("i: %v j: %v", i,j)
		name1 := *persons[i].Name
		name2 := *persons[j].Name
		return name1 < name2
	})

	for _, person := range persons {
		log.Printf("%v", *person.Name)
	}
}
