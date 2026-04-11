package main

import (
	"fmt"
)

func main() {
	const pi = 3.14159
	const greeting = "Hello, Go!"
	fmt.Println("Pi:", pi)
	fmt.Println("Greeting:", greeting)
	// Uncommenting the following lines will cause a compile-time error
	// because constants cannot be reassigned.
	// pi = 3.14 // this will cause an error: cannot assign to pi (constant)
	// greeting = "Hi!" // this will cause an error: cannot assign to greeting (constant)
}