package main

import "fmt"

func main() {
	nums := []int{0,1,2,3,4}
	newNums := nums[:2]
	newNums[1] = 100

	fmt.Printf("%v\n", nums[:2])
	fmt.Printf("%v\n", nums[0:])

	fmt.Printf("%v\n", nums)
}
