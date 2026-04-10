package main

import (
	"fmt"
	"maps"
)

func main() {
	// creating map

	// m := make(map[string]string)
	// m["name"] = "sujal"
	// m["lang"] = "golang"

	// fmt.Println(m["name"]) // sujal
	// fmt.Println(m["lang"]) // golang
	// IMP: if key is not present in map, it returns falsy value
	// fmt.Println(m["phone"]) // empty value
	// fmt.Println(m)
	/*
		m := make(map[string]int)
		m["age"] = 18
		m["price"] = 50
		m["num"] = 12
		fmt.Println(m["phone"]) // 0
		fmt.Println(m["age"]) // 18
		fmt.Println(len(m)) // 2
		delete(m, "num") // deletes num key
		fmt.Println(m) //map[age:18 price:50]
		clear(m) // deletes all items from map
		fmt.Println(m) // map[]
	*/

	// maps without make func
	/*
	   m := map[int] string {
	   	1: "one",
	   	2: "two",
	   	3: "three",
	   	4: "four",
	   }
	   // fmt.Println(m)
	*/
	// m := map[int]string{
	// 	1: "one",
	// 	2: "two",
	// 	3: "three",
	// 	4: "four",
	// }

	// v, ok := m[4]
	// if ok {
	// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("not ok")
	// }
	// fmt.Println(ok, v)

	m1 := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
	}
	m2 := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	fmt.Println(maps.Equal(m1, m2)) // false if not equal

}
