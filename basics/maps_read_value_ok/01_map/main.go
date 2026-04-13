package main

import "fmt"

func main() {
	// map is like object in javascript
	// map[keyType]valueType

	ages := map[string]int{
		"alice": 30,
		"bob":   25,
	}

	println(ages["alice"]) // 30
	fmt.Println(ages)

	var scores map[string]int
	fmt.Println(scores) // map[]
	scores = make(map[string]int)
	scores["math"] = 90
	scores["english"] = 85
	fmt.Println(scores) // map[english:85 math:90]
}
