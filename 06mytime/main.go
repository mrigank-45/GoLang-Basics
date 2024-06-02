package main

import (
	"fmt"
	"time"
)

func main() {

	// using time package
	var current = time.Now()
	fmt.Println("Current time is: ", current)

	// changing format
	fmt.Println("Current time is: ", current.Format("02-02-2006 Monday 15:04:05")) // date day time

	var createdDate = time.Date(2020, time.January, 1, 4, 3, 24, 50, time.UTC)
	fmt.Println("Created date is: ", createdDate)

}
