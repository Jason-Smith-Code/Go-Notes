// In each file we need to assign the code to a package, we do this with the package keyword, followed by a name. Main is the main package.
package main

// the import keyword allows us to import packages to use in our code. each package should be entered on a new line, no separators required
import (
	"fmt"
)

// a file almost always contains a function, the function declaration syntax is almost identical to javascript.
func main() {
	// here we access the fmt package, use the PrintIn method, and pass in a string as the arguement.
  fmt.Println("Hello World!")
}


//  declaring a variable : var variablename type = value


// Regular expressions