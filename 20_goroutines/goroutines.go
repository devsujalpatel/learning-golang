package main

import (
	"fmt"
)

// func task(id int) {
// 	fmt.Println("doing task", id)
// }

func main() {
	for i := 0; i <= 10; i++ {
		// go task(i)
	go func (i int) {
			fmt.Println(i)
		}(i)
	}


}
