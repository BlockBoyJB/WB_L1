package main

import (
	"fmt"
	"math"
)

func main() {
	p1 := NewPoint(1, 1)
	p2 := NewPoint(4, 5)

	fmt.Println(p1.Distance(p2))
}

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) Distance(o *Point) float64 {
	// по формуле: квадратный корень из суммы квадрата разностей координат осей
	return math.Sqrt(math.Pow(p.x-o.x, 2) + math.Pow(p.y-o.y, 2))
}
