// An example of using functionality from a package.
package main

import (
	"fmt"

	"github.com/EngineerKamesh/gofullstack/volume1/section2/greetingspackage"
)

func main() {

	greetingspackage.PrintGreetings()
	greetingspackage.GopherGreetings()

	fmt.Println("The value of the Magic Number is:", greetingspackage.MagicNumber)

}
