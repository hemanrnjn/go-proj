package main

import "fmt"

func anonymous() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	newInt := anonymous()
	fmt.Println(newInt())
	fmt.Println(newInt())
	fmt.Println(newInt())
}
