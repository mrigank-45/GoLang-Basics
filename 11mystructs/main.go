package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	var hitesh = User{"Hitesh", "hitesh@go.dev", true, 16}
	fmt.Println(hitesh)
	fmt.Printf("hitesh details are: %+v\n", hitesh) // %+v will print the field names as well
	fmt.Printf("Name is %v and email is %v.", hitesh.Name, hitesh.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
