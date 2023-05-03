package main

import (
	"log"
	"time"
)

type S struct {
	data string
}

func (s S) Read() string {
	return s.data
}

func (s *S) Write(data string) {
	s.data = data
}

func main() {
	d := S{"dicky novanto"}
	log.Printf("data being read: %v", d.Read())
	d.Write("asu")
	log.Printf("data now: %v", d.Read())

	ma := map[int]S{1: {"A"}}
	log.Printf("read from map: %v", ma[1].Read())
	(*S).Write(&d, "1233")
	log.Printf("now now: %v", d.Read())

	time.Now().AddDate()
}
