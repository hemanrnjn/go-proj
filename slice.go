package main

import "fmt"

func main() {

	s := make([]int, 3)
	fmt.Println("Slice: ", s)
	s = append(s, 2, 3)
	fmt.Println("Slice is: ", s)

}
