package main

import "fmt"

func main() {
	views := []int{100, 200, 150, 300}
	// for range

	total := 0
	for i, v := range views {
		fmt.Println("day", i, "views:", v)
		total += v
	}
	println("Total views:", total)
}
