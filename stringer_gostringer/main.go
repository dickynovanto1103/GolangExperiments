package main

import "log"

type AccessStatus int

const (
	_ AccessStatus = iota
	Denied
	Allowed
)

func (a AccessStatus) String() string {
	switch a {
	case Denied:
		return "Denied"
	case Allowed:
		return "Allowed"
	}
	return ""
}

func (a AccessStatus) GoString() string {
	return a.String()
}

func main() {
	tryFmtPrinting(Allowed)
	tryFmtPrinting(Denied)
}

func tryFmtPrinting(accessStatus AccessStatus) {
	log.Printf("accessStatus: %v", accessStatus)
	log.Printf("%%v: %v %d", accessStatus, accessStatus) // will print the result from String() if we call it with %v
	log.Printf("%%#v: %#v", accessStatus)                //it will print string if there is GoString() method, if no, it will print the number
	log.Printf("%%T: %T", accessStatus)
}
