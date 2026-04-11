package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Iteration %d\n", i)
	}
	// let print a pyramid of stars perfect pyramid
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}