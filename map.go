package main

import "fmt"

func main() {
	s := make(map[string]int)

	s["age"] = 22
	s["fav number"] = 7
	
	fmt.Println(s)
	delete(s, "fav number")
	_, prs := s["fav number"]
	fmt.Println(prs)
}
