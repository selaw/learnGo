package main

import (
	"encoding/json"
	"fmt"
)

var m map[string]string

func main() {
	var name string
	var add string
	m = make(map[string]string)
	fmt.Println("Enter Name:")
	fmt.Scan(&name)
	fmt.Println("Enter Address:")
	fmt.Scan(&add)
	m["Name"] = name
	m["Address"] = add
	fmt.Println(m)
	fmt.Println()
	fmt.Println("******Json Object*******")
	j, a := json.Marshal(m)
	fmt.Println(string(j), a)
}
