package main

import "fmt"

var temp = 10  // allowed outside function
// jwtToken := "some" // not allowed outside function

const Token = "some"  // capital first letter represents public variable

func main() {
	var username string = "mrigank"  // default "" for string
	fmt.Println("Hello, ", username, "!")

	var male bool = true
	fmt.Println(male, "!")

	var age int = 21  // default 0 for int
	fmt.Println(age, "!")

	var age1 float64 = 21.5545  // float is not valid
	fmt.Println(age1, "!")


	var name = "mrigank"  // type inference
	fmt.Println(name, "!")

	// no var style
	age2 := 21
	fmt.Println(age2, "!")
}
