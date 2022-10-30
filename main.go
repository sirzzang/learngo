package main

import (
	"fmt"
)

func main() {

	c := make(chan string)

	people := [5]string{"sirzzang", "eraser", "ieere", "ruby", "sir"}
	for _, person := range people {
		go isCool(person, c) // two goroutines
	}
	for i := 0; i < len(people); i++ {
		fmt.Print("waiting for: ", i, "\n")
		fmt.Println("received: " + <-c) // blocking operation
	}
}

func isCool(person string, c chan string) {
	c <- person + " is cool" // send true value to channel
}
