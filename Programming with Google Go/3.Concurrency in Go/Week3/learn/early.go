package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("New Normal")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Normal")
}
