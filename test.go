package main

import "fmt"
import "time"

func loop(val chan string) {
	for {
		select {
			case <-val: 
				fmt.Println("Congrats! You made out of it")
				return
			default:
				time.Sleep(1 * time.Second)
				fmt.Println("Dormammu. I have come to bargain")
		}
	}
}

func valInput(val chan string) {
	var input string
	for {
		fmt.Scanln(&input)
		if input=="Deal" || input=="deal" {
			val <- input
			break	
		}
	}
}

func main() {
	fmt.Println("Hint: You are stuck in a time loop. Enter 'Deal' to break out")
	val := make(chan string)
	go valInput(val)
	go loop(val)

	var inp string
	fmt.Scanln(&inp)
	fmt.Println("Entered value is: ", inp)
}
