package main

import (
	"fmt"
	"strings"
)

func findian() {
	var s string

	fmt.Println("Please enter a string")
	_, err := fmt.Scan(&s)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	s = strings.ToLower(s)

	if strings.Contains(s, "a") && strings.HasPrefix(s, "i") && strings.HasSuffix(s, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
