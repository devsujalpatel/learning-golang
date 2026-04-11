package main

// Iterating for data structures
func main() {
	// nums := []int{6, 7, 8}
	// sum := 0
	// for i, num := range nums {
	// 	println(i, num)
	// 	// sum = sum + num
	// }
	// println(sum)

	// m := map[string]string{"fname": "sujal", "lname": "patel"}
	// for k, v := range m {
	// 	println(k, v)
	// }

	// unicode code point rune
	// starting byte of rune
	// 255 -> 1 byte
	// 300 -> 2 bytes
	for i, c := range "golang" {
		println(i, string(c))
	}

}
