package main

import (
	"fmt"
	"sync"
)

var on sync.Once

func setup() {
	fmt.Println("Init")
}
func dostuff() {
	on.Do(setup)
	fmt.Println("hello")
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go dostuff()
	go dostuff()
	wg.Wait()
}
