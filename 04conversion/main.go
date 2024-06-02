package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your rating: ")
	rating, _ := reader.ReadString('\n')

	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64) // conversion from string to float64

	if err != nil {
		fmt.Println(err)
	} else {
		var newRating = numRating + 1
		fmt.Println("newRating is", newRating)
	}
}
