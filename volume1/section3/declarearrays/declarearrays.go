package main

import "fmt"

// 1st attempt at populating the integer array
func populateIntegerArray(input [5]int) {

	input[0] = 3
	input[1] = 6
	input[2] = 9
	input[3] = 12
	input[4] = 15

}

// 2nd attempt at populating the integer array
func populateIntegerArrayWithReturnValue(input [5]int) [5]int {

	input[0] = 3
	input[1] = 6
	input[2] = 9
	input[3] = 12
	input[4] = 15
	return input
}

func beatlesArrayExample() {

	// Declare and initialize values in an array
	var beatles [4]string
	beatles[0] = "John"
	beatles[1] = "Paul"
	beatles[2] = "Ringo"
	beatles[3] = "George"
	fmt.Printf("Beatles consists of: %v\n", beatles)
	fmt.Println("The name of third band member in the beatles array is", beatles[2])

	fmt.Println("Length of beatles: ", len(beatles))

	// In Go, arrays are values. When we assign one array to another, all the elements from
	// the array on the right hand side are copied over to the array on the left hand side.
	var greatRockBandFromThe60s [4]string
	greatRockBandFromThe60s = beatles
	fmt.Printf("Members from a great rock band from the 1960s: %v\n", greatRockBandFromThe60s)

	// Since arrays are values, equality comparisons of two arrays are done value for value.
	// Note that the beatles array and the greatRockBandFromThe60s array have two different
	// addresses in memory. The value comparison only checks the values of the arrays, and
	// the memory address of the two arrays is not a criteria for the comparison.
	fmt.Printf("beatles mem address: %p\n", &beatles)
	fmt.Printf("greatRockBandFromThe60s mem address: %p\n", &greatRockBandFromThe60s)
	if beatles == greatRockBandFromThe60s {
		fmt.Println("The beatles array equals the greatRockBandFromThe60s array")
	}
}

func u2ArrayExample() {

	// Declare and initialize using the := operator. Instead of writing 4 lines of code
	// to initialize the array, we get the job done in 1 line of code using an array
	// literal value
	u2 := [4]string{"Bono", "Edge", "Adam", "Larry"}
	fmt.Printf("U2 consists of: %v\n", u2)
	fmt.Println("The name of second band member in the u2 array is", u2[1])

	fmt.Println("Length of u2: ", len(u2))
}

func main() {

	// Declare an array of 5 integers
	// Note: If we do not initialize a value for the array, they will default to the
	// zero value of the type. In the case of integers, we expect the zero value to be 0.
	var myArray [5]int
	fmt.Printf("Contents of myArray: %v\n\n", myArray)

	// Arrays are passed by value to functions, meaning a copy of the array is passed
	// and not a reference (pointer) to the array.
	populateIntegerArray(myArray)
	fmt.Printf("Contents of myArray: %v\n\n", myArray)

	myArray = populateIntegerArrayWithReturnValue(myArray)
	fmt.Printf("Contents of myArray: %v\n\n", myArray)

	// Use the built in len function to get the length of an array
	fmt.Println("Length of myArray: ", len(myArray))

	beatlesArrayExample()

	u2ArrayExample()

}
