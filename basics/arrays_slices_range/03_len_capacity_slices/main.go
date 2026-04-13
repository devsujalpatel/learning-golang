package main

import "fmt"

func main() {
	// length and capacity of slices
	scores := make([]int, 0, 5)
	fmt.Println(scores, len(scores), cap(scores)) // Output: [] 0 5

	// appending to the slice
	scores = append(scores, 90, 80, 70)
	fmt.Println(scores, len(scores), cap(scores)) // Output: [90 80 70] 3 5

	// appending more elements than capacity it will automatically increase the capacity (usually doubling it)
	scores = append(scores, 60, 50, 40)
	fmt.Println(scores, len(scores), cap(scores)) // Output: [90 80 70 60 50 40] 6 10 (capacity doubled)

	todos := []string{"Buy groceries", "Clean the house", "Pay bills"}

	more := []string{"Go for a walk", "Read a book"}
	todos = append(todos, more...)
	fmt.Println(todos) // Output: [Buy groceries Clean the house Pay bills Go for a walk Read a book]
}
