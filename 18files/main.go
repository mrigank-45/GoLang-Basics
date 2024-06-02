package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file"

	file, err := os.Create("./mylcogofile.txt") // creating a file

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content) // writing to file
	checkNilErr(err)
	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./mylcogofile.txt")
}

func readFile(filname string) {
	databyte, err := ioutil.ReadFile(filname) // reading from file
	checkNilErr(err)

	fmt.Println("Text data inside the file is \n", string(databyte))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
