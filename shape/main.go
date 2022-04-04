package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {
	s := square{sideLength: 2.5}
	t := triangle{height: 3, base: 6}

	printArea(s)
	printArea(t)
}

func printArea(s shape) {
	fmt.Println("Shape area:", s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
