package main

import "fmt"

func main() {

	// We create a buffered channel of strings with a capacity of 3
	// This means the channel buffer can hold up to 3 values
	messageQueue := make(chan string, 3)
	messageQueue <- "one"
	messageQueue <- "two"
	messageQueue <- "three"

	// We drain the messageQueue by receiving all the values from the
	// buffered channel.
	fmt.Println(<-messageQueue)
	fmt.Println(<-messageQueue)
	fmt.Println(<-messageQueue)

}
