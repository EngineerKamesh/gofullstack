package main

import (
	"fmt"
	"sync"
)

var greetings string
var howdyDone chan bool
var mutex = &sync.Mutex{}

func howdyGreetings() {

	mutex.Lock()
	greetings = "Howdy Gopher!"
	mutex.Unlock()
	howdyDone <- true
}

func main() {

	howdyDone = make(chan bool, 1)
	go howdyGreetings()
	mutex.Lock()
	greetings = "Hello Gopher!"
	fmt.Println(greetings)
	mutex.Unlock()
	<-howdyDone
}
