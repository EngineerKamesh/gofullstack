package main

import "fmt"

func main() {

	// We create a buffered channel of strings with a capacity of 3
	// This means the channel buffer can hold up to 3 values
	messageQueue := make(chan string, 3)
	messageQueue <- "one"
	messageQueue <- "two"
	messageQueue <- "three"

	// Even though we close this non-empty channel, we can still receive
	// the remaining values (see below)
	close(messageQueue)

	// We use the range keyword to iterate over each element as it gets
	// received from the messageQueue.
	for m := range messageQueue {
		fmt.Println(m)
	}

}
