package main

import (
	"github.com/gogo/protobuf/proto"
	"log"
	"sort"
)

type Person struct {
	Age *uint64
}

func main() {
	persons := []*Person{
		{
			Age: proto.Uint64(1),
		},
		{
			Age: proto.Uint64(123),
		},
	}

	mapper := make(map[int]*Person)
	mapper[1] = &Person{Age: proto.Uint64(1),}
	log.Printf("mapper 2: %v", mapper[2])
	persons = append(persons, mapper[2])

	for _, person := range persons {
		log.Printf("%v", *person.Age)
	}

	sort.Slice(persons, func(i, j int) bool {
		log.Printf("i: %v j: %v", i,j)
		name1 := *persons[i].Age
		name2 := *persons[j].Age
		return name1 < name2
	})

	for _, person := range persons {
		log.Printf("%v", *person.Age)
	}
}
