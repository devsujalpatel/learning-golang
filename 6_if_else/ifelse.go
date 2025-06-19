package main

import "fmt"

func main() {
	// age := 14
	// if age >= 18 {
	// 	fmt.Println("You are an adult")
	// } else {
	// 	fmt.Println("You are not an adult")
	// }
	// age := 18
	
	// if age >= 18 {
	// 	fmt.Println("You are an adult")
	// } else if age >= 12 {
	// 	fmt.Println("You are a teenager")
	// } else {
	// 	fmt.Println("You are not an adult")
	// }

	var role = "admin"
	var hasPermission = false

	if role == "admin" && hasPermission {
		fmt.Println("You are an admin")
	} else {
		fmt.Println("Permission false")
	}

	if role == "admin" || hasPermission {
		fmt.Println("You are an admin")
	} else {
		fmt.Println("Permission false and you are not admin")
	}
	

}
