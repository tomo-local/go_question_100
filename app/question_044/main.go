package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func printArea(s Shape) {
	fmt.Printf("%T of Area: %0.2f\n", s, s.Area())
}


func main() {
	rect := Rectangle{width: 10, height:5}
	circ := Circle{radius: 4}

	printArea(rect)
	printArea(circ)
}
