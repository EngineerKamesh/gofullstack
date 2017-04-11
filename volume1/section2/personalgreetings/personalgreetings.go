package main

import (
	"flag"
	"fmt"
)

func main() {

	var gopherName string
	flag.StringVar(&gopherName, "gophername", "Gopher", "The name of the Gopher")
	flag.Parse()
	fmt.Println("Hello " + gopherName + "!")
}
