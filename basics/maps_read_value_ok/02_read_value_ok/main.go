package main

import "fmt"

func main() {
	points := map[string]int{
		"a": 10,
		"b": 0, // valid value
	}

	fmt.Println("a", points["a"])
	fmt.Println("b", points["b"])
	fmt.Println("c", points["c"]) // zero value for int is 0

	valB, okB := points["b"]
	fmt.Println(valB, okB) // b 0 true

	valC, okC := points["c"]
	fmt.Println(valC, okC) // c 0 false

	if val, ok := points["c"]; ok {
		fmt.Println("c", val)
	} else {
		fmt.Println("c does not exist")
	}

	prices := map[string]int{
		"xyz": 100,
		"def": 1800,
	}

	total := 0

	for item, price := range prices {
		fmt.Printf("Item: %s, Price: %d\n", item, price)
		total += price
	}
	fmt.Printf("Total: %d\n", total)

	for item := range prices { // only keys
		fmt.Printf("Item: %s\n", item)
	}

	for _, price := range prices { // only values
		fmt.Printf("Price: %d\n", price)
	}
}
