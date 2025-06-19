package main

import (
	"fmt"
	"time"
)

func main () {
	// simple switch
	// i := 2
// does'nt need break
	// switch i {
	// case 1:
	// 	fmt.Println(1)
	// case 2:
	// 	fmt.Println(2)
	// case 3:
	// 	fmt.Println(3)
	// default: // optional
	// 	fmt.Println("other")
	// }

	// Multiple condition switch

	switch time.Now().Weekday() {
		case time.Saturday, time.Sunday:
			fmt.Println("It's the weekend")
		default:
			fmt.Println("It's a workday")
	}
	
}
