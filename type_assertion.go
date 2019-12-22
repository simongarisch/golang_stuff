package main

import "fmt"

type Shape interface {
	Area() float64
}

type Square struct{}

type Rectangle struct{}

func (s Square) Area() float64 {
	return 10.0
}

func (r Rectangle) Area() float64 {
	return 20.0
}

func PrintArea(s Shape) {
	fmt.Println(s.Area())
}

func main() {
	square := Square{}
	rectangle := Rectangle{}
	shapes := []Shape{square, rectangle}

	for _, shape := range shapes {
		PrintArea(shape) // 10, 20
	}
	fmt.Println("----------")

	var shape Shape
	shape = square

	rec, ok := shape.(Rectangle) // is this shape a rectangle
	fmt.Println(rec, ok)         // {} false
	if ok {
		fmt.Println("We have a rectangle...")
	} else {
		fmt.Println("This is not a rectangle.")
	}

	// or we could use a switch
	switch shape.(type) {
	case Rectangle:
		fmt.Println("We have a rectangle...")
	case Square:
		fmt.Println("We have a square...")
	}
}
