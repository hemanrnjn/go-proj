package main

import "fmt"

func f(str string) {
	for i := 0; i < 10; i++ {
		fmt.Println(str, " : ", i) 
	}
}

func main() {
	f("main")

	go f("thread")

	go func(){
		fmt.Println("Interrupt")
	}()

	fmt.Scanln()
	fmt.Println("done")
	
}
