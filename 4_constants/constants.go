package main

import "fmt"

const age = 30 // possible outside main

// const := "hello" // not possible outside main

func main() {
	const name string = "golang"

	// name = "javascript" wrong // cannot assign to name
	fmt.Println(name, age)
}
