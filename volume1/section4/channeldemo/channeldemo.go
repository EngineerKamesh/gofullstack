package main

import "fmt"

func writeMessageToChannel(message chan string) {

	message <- "Hello Gopher!"

}

func main() {

	fmt.Println("Channel Demo")

	message := make(chan string)
	go writeMessageToChannel(message)

	fmt.Println("Greeting from the message channel: ", <-message)

	close(message)
}
