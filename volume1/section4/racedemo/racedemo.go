// An example of a data race condition.
// Execute the program like so: "go run -race racedemo.go" to detect the race condition.
package main

import "fmt"

var greetings string
var howdyDone chan bool

func howdyGreetings() {

	greetings = "Howdy Gopher!"
	howdyDone <- true
}

func main() {

	howdyDone = make(chan bool, 1)
	go howdyGreetings()
	greetings = "Hello Gopher!"
	fmt.Println(greetings)
	<-howdyDone
}
