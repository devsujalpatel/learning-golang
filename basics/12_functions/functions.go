package main

// func add(a int, b int) int {
// 	return (a + b)
// }

// func getLanguages() (string, string, string, bool) {
// 	return "golang", "javascript", "c", true
// }

//	func processIt(fn func(a int) int) {
//		fn(2)
//	}
func processIt() func(a int) int {
	return func(a int) int {
		return 5
	}
}

func main() {
	// result := add(2, 6)
	// println(result)
	// println(getLanguages())
	// lang1, lang2, lang3, _ := getLanguages()
	// println(lang1, lang2, lang3)
	// fn := func(a int) int {
	// 	return a
	// }
	// println(fn)
	fn := processIt()
	println(fn(2))
}
