package main

import (
	"fmt"
)

var done chan bool = make(chan bool)

func printGreetings(source string) {

	for i := 0; i < 9; i++ {
		fmt.Println("Hello Gopher!", i, source)
	}

	if source == "goroutine" {
		done <- true
	}

}

func main() {

	go printGreetings("goroutine")
	printGreetings("main function")

	<-done
}
