package main

import "fmt"

type Shape interface {
	Draw()
}

type Circle struct{}
type Square struct{}
type Triangle struct{}

type Drawing struct {
	Shapes []Shape
}

func NewDrawing() Drawing {
	return Drawing{[]Shape{}}
}

func (drawing *Drawing) AddShape(shape Shape) {
	drawing.Shapes = append(drawing.Shapes, shape)
}

func (drawing *Drawing) Draw() {
	for _, shape := range drawing.Shapes {
		shape.Draw()
	}
}

func (shape Circle) Draw() {
	fmt.Println("Drawing a circle")
}

func (shape Square) Draw() {
	fmt.Println("Drawing a square")
}

func (shape Triangle) Draw() {
	fmt.Println("Drawing a triangle")
}

func main() {
	drawing := NewDrawing()
	for _, shape := range []Shape{Circle{}, Square{}, Triangle{}} {
		drawing.AddShape(shape)
	}
	fmt.Println(drawing) // {[{} {} {}]}

	drawing.Draw() // calls draw method for each of drawing's components
	// Drawing a circle
	// Drawing a square
	// Drawing a triangle
}
