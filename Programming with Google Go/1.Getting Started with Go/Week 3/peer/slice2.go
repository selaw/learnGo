package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	// Variable for the input
	var input string

	// Make a slice with length 0 and capacity 5
	numbers := make([]int, 0, 3)

	// Prompt user
	fmt.Println("Please enter one or more inergers. Input X to terminate")

	// Loop forever
	for {
		// Read input
		fmt.Scan(&input)
		if input[0] == 'X' {
			break
		}

		// Attempt to parse the input as an integer
		inputInt, e := strconv.Atoi(input)

		// If an error was thrown, notify the user and do not enter parsed value in result
		if e != nil {
			fmt.Println("Please enter only integers, or X to terminate")
			continue
		}

		// Add the integer to the result
		numbers = append(numbers, inputInt)

		// Sort the result
		sort.Ints(numbers)

		// Print the result
		fmt.Println(numbers)
	}
}
