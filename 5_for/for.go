package main

import "fmt"

// for loop -> only construct in go for looping
func main () {
	// while loop
	// i := 1
	// for i <= 5 {
	// 	println(i)
	// 	i++
	// }
	// Infinite while loop
	// for {
	// 	i++
	// 	println("hello", i)
	// }

	// Classic for loop

	for i := 0; i <= 3; i++ {
		// break
		if i == 2 {
			break
		}
		// continue 
		// if i == 2 {
		// 	continue
		// }
			fmt.Println(i)
		
	}



}
