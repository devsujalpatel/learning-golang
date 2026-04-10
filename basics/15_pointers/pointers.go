package main


// by value
// func changeNum(num int) {
// 	num = 5
// 	println("In changeNum", num)
// }

// by reference
func changeNum(num *int) {
	*num = 5
	println("In changeNum", *num)
}

func main() {
	num := 1
	println(num)
	changeNum(&num)
	// println("Memory address", &num)
	println("After changeNum in main", num)
}
