package main

import "fmt"

func main() {
	var input string
	val, err:= fmt.Scanln(&input)
	fmt.Println("Input: ", val, "or :", err)
}
