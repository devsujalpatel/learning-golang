package main

import (
	"fmt"
	"sort"
)

func main() {
	// common collection type
	// dynamic and can grow or shrink in size
	// []type{...}

	names := []string{"Alice", "Bob", "Charlie"} // Slice literal

	fmt.Println(names[len(names)-1]) // Accessing the last element of the slice

	// appending to a slice
	names = append(names, "David")
	fmt.Println(names)

	// slicing a slice
	subNames := names[1:3] // Slicing from index 1 to 2 (exclusive)
	fmt.Println(subNames)
	fmt.Println(names) // Original slice remains unchanged

	var nums []int    // nil slice, can be appended to
	fmt.Println(nums) // Output: []
	nums = append(nums, 1, 4, 3)
	fmt.Println(nums) // Output: [1 4 3]
	sort.Ints(nums) // Sorting the slice
	fmt.Println(nums) // Output: [1 3 4]
}
