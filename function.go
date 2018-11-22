package main

import "fmt"

func add(a int, b int) int {
	return a+b
}

func multi() (int, int) {
	return 1,2
}

func main() {
	res := add(1, 2)
	fmt.Println("Sum of 1 and 2 is: ", res)

	multiRes1, multiRes2 := multi()
	fmt.Println("Multiple values of multi function is: ", multiRes1, " and ", multiRes2)
}
