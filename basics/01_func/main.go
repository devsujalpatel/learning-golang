package main

import "fmt"

func addTwo(a int, b int) int {
	return a + b
}

func SumAndProduct(a int, b int) (int, int) {
	sum := a + b
	product := a * b
	return sum, product
}

func main() {
	fmt.Println(addTwo(3, 5))
	fmt.Println(SumAndProduct(3, 8))
}
