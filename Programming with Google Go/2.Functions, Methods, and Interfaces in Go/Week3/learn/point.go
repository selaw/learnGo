package main

import (
	"fmt"
)

func (p *Point) OffsetX(v float64) {
	p.x = p.x + v
}

func main() {
	p := Point{3, 4}
	p.OffsetX(5)
	fmt.Println(p.x)
}