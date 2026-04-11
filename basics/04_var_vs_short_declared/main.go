package main

import (
	"fmt"
)

func main() {
	var city string = "new york" // using var keyword to declare a variable with explicit type.
	fmt.Println("City:", city)
	
	country := "usa" // using short variable declaration, the type of country is inferred as string.
	fmt.Println("Country:", country)
	likes, comments := 100, 20 // multiple short variable declaration for likes and comments.
	fmt.Println("Likes:", likes, "Comments:", comments)
}