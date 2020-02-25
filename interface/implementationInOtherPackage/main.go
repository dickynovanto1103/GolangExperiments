package main

import "github.com/dickynovanto1103/GolangExperiments/interface/implementationInOtherPackage/sampleImplementation"

func main() {
	anyStruct := &sampleImplementation.AnyStruct{}
	st := sampleImplementation.NewST(anyStruct)
	st.Process()
}
