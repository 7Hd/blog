package main

import (
	"fmt"
)

type Point struct {
	x float64
	y float64
}

func (p Point) getXAddOne() float64 {
	return p.x + 1
}

func main() {
	p := Point{x: 3.0, y: 4.0}

	fmt.Println(p.x)
	fmt.Println(p.y)
	fmt.Println(p.getXAddOne())
}
