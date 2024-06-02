package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var username string
	// fmt.Println("Enter your name: ")
	// fmt.Scanln(&username)
	// fmt.Println("Hello, ", username, "!")

	// **Using bufio
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")

	// comma ok or comma err syntax
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello, ", name)

}
