package main

import "fmt"

func variad(nums ...int) int {
	fmt.Println(nums, " ")
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	fmt.Println(variad(1,2,3,4))
	fmt.Println(variad(1,2,3,4,5,6,7))
}
