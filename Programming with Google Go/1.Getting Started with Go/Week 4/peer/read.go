package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {
	people := make([]Person, 0, 100)
	var filePath string
	fmt.Println("Enter the path to the file")
	fmt.Scan(&filePath)
	file, _ := os.Open(filePath)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		names := strings.Split(scanner.Text(), " ")
		person := Person{fname: names[0], lname: names[1]}
		people = append(people, person)
	}
	fmt.Println(people)
	file.Close()
}
