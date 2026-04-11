package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Alice"
	lowercase := "hello, world!"
	fmt.Println("This will be uppercase:", strings.ToUpper(lowercase))
	fmt.Println("Name:", name)
}