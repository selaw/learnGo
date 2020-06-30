package main

import (
	"fmt"
	"sync"
)

func inc(wg *sync.WaitGroup, i *int) {
	*i = *i + 1
	fmt.Println(*i)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var i = 0

	wg.Add(2)
	go inc(&wg, &i)
	go inc(&wg, &i)
	wg.Wait()
}
