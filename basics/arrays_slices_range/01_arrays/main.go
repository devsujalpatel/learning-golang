package main

import "fmt"

func main() {
	var marks [3]int // fixed and cannot grow
	marks[0] = 90
	marks[1] = 80
	marks[2] = 70

	fmt.Println(marks)

	// array literal
	ages := [5]int{10, 20, 30, 40, 50}
	fmt.Println(ages)

	// array of strings
	names := [3]string{"Alice", "Bob", "Charlie"}
	fmt.Println(names)
	// length of array
	fmt.Println("Length of marks array:", len(marks))

}