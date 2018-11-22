package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	fmt.Println(person{"Himanshu Ranjan", 20})
	s := person{"Ashwani", 23}
	fmt.Println(s.name)
	s.name = "Suraj"
	fmt.Println(s.name)
}
