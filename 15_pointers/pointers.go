package main

func changeNum(num int) {
	num = 5
	println("In changeNum", num)
}

func main() {
	num := 1
	changeNum(num)

	println("After changeNum in main", num)
}
