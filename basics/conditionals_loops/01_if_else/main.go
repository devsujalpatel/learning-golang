package main

import (
	"fmt"
)

func main() {
	isActive := true
	if isActive {
		fmt.Println("User is active.")
	} else {
		fmt.Println("User is not active.")
	}

	age := 20
	if age < 18 {
		fmt.Println("User is a minor.")
	} else if age >= 18 && age < 65 {
		fmt.Println("User is an adult.")
	} else {
		fmt.Println("User is a senior.")
	}
}