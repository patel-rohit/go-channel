package main

import (
	"fmt"
	"time"
)

// Function
func myfun(mychnl chan string) {

	for v := 0; v < 10; v++ {
		mychnl <- fmt.Sprintf("GeeksforGeeks %v", v)
	}
	close(mychnl)
}

// Main function
func main() {

	now := time.Now()

	fmt.Println("The current datetime is:", now)

	// Creating a channel
	c := make(chan string)

	// calling Goroutine
	go myfun(c)

	// When the value of ok is
	// set to true means the
	// channel is open and it
	// can send or receive data
	// When the value of ok is set to
	// false means the channel is closed
	for {
		res, ok := <-c
		if ok == false {
			fmt.Println("Channel Close ", ok)
			break
		}
		fmt.Println("Channel Open ", res, ok)
	}
}
