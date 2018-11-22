package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sum: ", sum)

	m := map[string]int{"age": 22, "fav number":7}
	i := 1
	for k, val := range m {
		fmt.Println(i, ". ", k, " is: ", val)
		i++
	}
}
