package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func Distance(t1, t2 Point) float64 {
	return math.Sqrt(math.Pow(t1.x-t2.x, 2) + math.Pow(t1.y-t2.y, 2))
}

func main() {
	a := NewPoint(1, 1)
	b := NewPoint(2, 3)
	fmt.Println(Distance(a, b))
}
