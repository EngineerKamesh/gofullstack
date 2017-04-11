package main

import "fmt"

func main() {

	x := 9

	// An if statement
	if x == 9 {
		fmt.Println("x is equal to 9")
	}

	// An if statement followed by an else statement example
	if y := 18; y == 18 {
		fmt.Println("y is equal to 18")
	} else {
		fmt.Println("y is not equal to 18")
	}

	z := 36
	// An if - else if - else example
	if z == 1 {
		fmt.Println("z is equal to 1")

	} else if z == 2 {
		fmt.Println("z is equal to 2")

	} else if z == 3 {
		fmt.Println("z is equal to 3")

	} else {
		fmt.Println("z does not equal 1, 2, or 3")
	}

}
