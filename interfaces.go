package main

import (
	"fmt"
	"math"
)

// Shape is a shape interface
type Shape interface {
	Area() float64
}

type Square struct {
	Length float64
}

func (s *Square) Area() float64 {
	return s.Length * s.Length
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func main() {
	square := &Square{10}
	circle := &Circle{5}

	fmt.Println(square.Area())
	fmt.Println(circle.Area())

	shapes := []Shape{square, circle}
	fmt.Println(shapes)
	for _, shape := range shapes {
		fmt.Println(shape.Area())
	}
}
