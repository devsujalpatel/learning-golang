package main

import "fmt"


// func printSlice[T int | string #or comparable ](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

type stack[T int | string] struct {
	elements []T
}

func main() {

	myStack := stack[int] {
		elements: []int{1, 2, 3},
	}
	fmt.Println(myStack)
	myStack2 := stack[string] {
		elements: []string{"sujal", "patel", "golang", "typescript"},
	}
	fmt.Println(myStack2)

	// nums := []int{1, 2, 3}
	// names := []string{"sujal", "patel", "golang", "typescript"}
	// printSlice(names)
	// printSlice(nums)
}
