package main

import (
	// "fmt"

	// "github.com/devsujalpatel/podcast/auth"
	"github.com/devsujalpatel/podcast/user"
	"github.com/fatih/color"
)

func main() {
	// auth.LoginWithCredentials("sujal", "sujal123")

	// session := auth.GetSession()

	// fmt.Println(session)

	user := user.User{Name: "Sujal", Age: "18"}
	// fmt.Println(user.GetName())
	// fmt.Println(user.GetAge())

	color.Red("Hello World!")
	color.Blue(user.GetName())
	age := user.GetAge()
	color.HiGreen(age)

}
