package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.length * s.width
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func info(s shape) {
	fmt.Printf("The area of your shape is: %.2f\n", s.area())
}

func main() {
	square1 := square{
		length: 100,
		width:  100,
	}

	circle1 := circle{
		radius: 100,
	}

	info(square1)
	info(circle1)
}
