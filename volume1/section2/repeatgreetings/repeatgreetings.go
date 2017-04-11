package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	var repeatCount int
	var err error

	if len(os.Args) >= 2 {
		repeatCount, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < repeatCount; i++ {
			fmt.Println("Hello Gopher!")
		}

	}

}
