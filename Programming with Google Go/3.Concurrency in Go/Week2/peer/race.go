package main

import "fmt"

var quit chan int
var glo int

func test() {
	fmt.Println(glo)
}

//The goroutine runs concurrently with the for loop and there's no synchronization performed.
//So race condition occurs
// To avoid race condition should synchronize the concurrent execution.
// If global variable is accessed, it should be protected with mutex.
func main() {
	glo = 0
	n := 10000
	quit = make(chan int, n)
	go test()
	for {
		quit <- 1
		glo++
	}
}
