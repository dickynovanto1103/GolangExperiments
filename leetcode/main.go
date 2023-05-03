package main

import "strings"

func haveConflict(event1 []string, event2 []string) bool {
    start1 := event1[0]
	end1 := event1[1]

	time1 := strings.Split(start1, ":")
	hh1 := time1[0]

	start2 := event2[0]
	end2 := event2[1]


}

func main() {
	
}