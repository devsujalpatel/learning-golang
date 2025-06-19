package main

import "fmt"

// for loop -> only construct in go for looping
func main () {
	// while loop
	i := 1
	for i <= 5 {
		println(i)
		i++
	}
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

	// Range loop

	for i := range 3 { // excludes 3 cause 0 to 3 meaning except 3
		fmt.Println(i)
	}


}
