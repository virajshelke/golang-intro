// we declare the package name here
package main

// The imports
import (
	"fmt"
)

// In go lang main function is the entry point
func main() {
	// Println is the function used to put data on stdout i.e. console
	// Note that the function name starts with a capital letter 'P' & that's because
	// in go lang we can only export (make things public) if we mark the symbol with starting
	// letter as captial (More on this in latter examples!)

	// Go lang doesn't need semi colons! Although the language is defined with semi-colons
	// The lexer in the compiler does this job for us so we can skip semi-colons
	fmt.Println("Hello world! It's my first go program")
}
