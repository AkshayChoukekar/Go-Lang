package main

import "fmt"

type shape interface {
	getArea() float64
}
type triangle struct {
	base   float64
	height float64
}
type square struct {
	sideLength float64
}

func main() {
	t := triangle{base: 5, height: 7}
	s := square{sideLength: 10}

	printArea(s)
	printArea(t)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return (1 / 2) * t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
