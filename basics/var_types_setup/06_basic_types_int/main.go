package main

import (
	"fmt"
)

func main() {
	// integers
	likes := 10
	likes++
	likes += 5 / 2
	// floats
	rating := 4.5
	rating *= 1.1
	fmt.Println("Likes:", likes)
	fmt.Println("Rating:", rating)
}