package main

import "fmt"
import "time"

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c<- "one"
	}()
	go func() {
		time.Sleep(1 * time.Second)
	}()
}
