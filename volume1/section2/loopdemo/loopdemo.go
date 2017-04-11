package main

import (
	"fmt"
	"time"
)

func main() {

	// Classic for loop
	for i := 0; i < 10; i++ {

		if i == 0 {
			continue
		}

		fmt.Println("Inside classic for loop, value of i is:", i)

	}
	fmt.Println("\n\n")

	// Single condition for loop
	j := -20
	for j != 0 {
		fmt.Println("Inside single condition loop, value of j is:", j)
		j++
	}
	fmt.Println("\n\n")

	loopTimer := time.NewTimer(time.Second * 9)
	// An infinite loop, don't worry we'll break out of it in 9 seconds
	for {

		fmt.Println("Inside the infinite loop!")

		<-loopTimer.C
		break

	}

}
