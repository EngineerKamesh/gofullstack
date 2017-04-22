package main

import (
	"fmt"
	"time"
)

func printGreetings(source string) {

	for i := 0; i < 9; i++ {
		fmt.Println("Hello Gopher!", i, source)
	}
	time.Sleep(time.Millisecond * 1)
}

func main() {

	go printGreetings("goroutine")
	printGreetings("main function")

}
