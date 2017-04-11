package greetingspackage

import "fmt"

// We indicate to Go that we want to export a function by upper casing the function's
// first letter.
func PrintGreetings() {

	fmt.Println("I'm priting a message from the PrintGreetings() function!")
}

// This function is unexported (since it has a lowercase first letter in the function name)
// Since it's unexported it will only be visible to functions that are within the greetingspackage
func printGreetingsUnexported() {

	fmt.Println("I'm printing a message from the printGreetingsUnexported() function!")

}
