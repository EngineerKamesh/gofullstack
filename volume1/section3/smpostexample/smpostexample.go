package main

import (
	"fmt"

	"github.com/EngineerKamesh/gofullstack/volume1/section3/socialmedia"
)

func main() {

	myPost := socialmedia.NewPost("EngineerKamesh", socialmedia.Moods["thrilled"], "Go is awesome!", "Check out the Go web site!", "https://golang.org", "", "", []string{"go", "golang", "programming language"})
	fmt.Printf("myPost: %+v\n", myPost)
}
