package main

import "fmt"

func main() {
	var a[5] int
	b := [5] int{1, 2, 3, 4, 5}
	fmt.Println("Array a is: ", a)
	fmt.Println("Array b is: ", b)

	var twoD[5][5] int
	for i := 0; i<5; i++ {
		for j := 0; j<5; j++ {
			twoD[i][j] = i+j
		}
	}
	fmt.Println("Two Dimensional array ex: ", twoD)
}
