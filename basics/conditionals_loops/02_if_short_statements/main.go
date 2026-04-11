package main

import "fmt"

func main() {
	items := 3
	pricePerItem := 39

	if total := items * pricePerItem; total > 100 {
		fmt.Printf("Total cost is %d, which is above the budget.\n", total)
	} else {
		fmt.Printf("Total cost is %d, which is within the budget.\n", total)
	}
}
