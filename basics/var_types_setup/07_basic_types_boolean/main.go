package main

import (
	"fmt"
)

func main() {
	isActive := true
	isComplete := false
	fmt.Println("Is Active:", isActive)
	fmt.Println("Is Complete:", isComplete)

	isBoth := isActive && isComplete
	isEither := isActive || isComplete
	fmt.Println("Is Both Active and Complete:", isBoth)
	fmt.Println("Is Either Active or Complete:", isEither)

}