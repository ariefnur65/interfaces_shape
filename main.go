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
	squareOne := square{40}
	triangleOne := triangle{height: 40,
		base: 20}
	printArea(squareOne)
	printArea(triangleOne)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println("Area ", s.getArea())
}
