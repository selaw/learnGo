package main

import (
	"fmt"
)

func apply(afungsi func(int) int, val int) int {
	return afungsi(val)
}

func increment(x int) int { return x + 1 }
func decrement(x int) int { return x - 1 }

func main() {
	fmt.Println(apply(increment, 2))
	fmt.Println(apply(decrement, 2))
}
