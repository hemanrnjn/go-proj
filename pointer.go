package main

import "fmt"

func setZero(a int) {
	a = 0
}

func setPtrZero(a *int) {
	*a = 0
}

func main() {
	a := 1
	b := &a
	setZero(a)
	fmt.Println(a)
	setPtrZero(b)
	fmt.Println(a)
}
