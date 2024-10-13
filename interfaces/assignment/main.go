package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}
type shape interface {
	getArea() float64
}

func main() {
	tri := triangle{
		height: 4.0,
		base:   2.0,
	}
	sq := square{
		sideLength: 4.0,
	}

	printArea(tri)
	printArea(sq)
}

func (tri triangle) getArea() float64 {
	return 0.5 * tri.base * tri.height
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
