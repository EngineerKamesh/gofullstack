package greetingspackage

import "fmt"

var MagicNumber int

// This is an exported function and can be called outside the package
func GopherGreetings() {

	fmt.Println("A very jolly hello my fellow gophers! I'm printing from the GopherGreetings() function")
	// Now we're calling an unexported package, because this function
	// is within the same package, we have access to call it
	printGreetingsUnexported()

}

// Example of a Packages init() function
func init() {

	MagicNumber = 108

}
