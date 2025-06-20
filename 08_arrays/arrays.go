package main

import "fmt"


func main () {
	// numbered sequence of specific length and type
	// zeroed values are falsy 
	// incase of string, empty string is falsy
	// incase of numbers, zero is falsy
	// incase of bool, false is falsy

	var num [4]int

	num[0] = 1
	num[1] = 2
	num[2] = 3
	num[3] = 4

	var vals [4]bool
	vals[2] = true
	fmt.Println(vals)

	fmt.Println(num[0])
	fmt.Println(num)
	// fmt.Println(len(num))

	var name [3]string
	name[0] = "go"
	fmt.Println(name)

	// assigning while declaring
	nums := [4]int{1, 2, 3, 4}
	fmt.Println(nums)

	// 2d arrays 
	TwoDarray := [2][2]int{{1, 2},{3, 4}}

	fmt.Println(TwoDarray)


	// use arrays if 
 // Fixed size, that is predictable 
 // Memory optimization
 // Constant time access


}