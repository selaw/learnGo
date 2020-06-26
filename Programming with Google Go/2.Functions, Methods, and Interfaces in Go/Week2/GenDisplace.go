package main

import (
	"fmt"
)

func GenDisplaceFn(acceleration, initialVelocity, initialDisplacement float64) func(time float64) float64 {
	return func(time float64) float64 {
		return 0.5*acceleration*time*time + initialVelocity*time + initialDisplacement
	}
}

func main() {

	var n int
	var acc float64
	n = 0
	for n == 0 {
		fmt.Println("Please input an acceleration: ")
		n, _ = fmt.Scanln(&acc)
	}
	var velo float64
	n = 0
	for n == 0 {
		fmt.Println("Please input an initial velocity: ")
		n, _ = fmt.Scanln(&velo)
	}
	var dis float64
	n = 0
	for n == 0 {
		fmt.Println("Please input an initial displacement: ")
		n, _ = fmt.Scanln(&dis)
	}
	var time float64
	n = 0
	for n == 0 {
		fmt.Println("Please input a time: ")
		n, _ = fmt.Scanln(&time)
	}

	fn := GenDisplaceFn(acc, velo, dis)
	fmt.Println("Result = ", fn(time))
}
