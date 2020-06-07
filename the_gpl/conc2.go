package main

import "fmt"

type Shape interface {
	Area() int
}

type Square struct {
	width int
}

type Rectangle struct {
	width  int
	height int
}

func (s Square) Area() int {
	return s.width * s.width
}

func (r Rectangle) Area() int {
	return r.width * r.height
}

func recordArea(shape Shape, areas map[Shape]int, ch chan bool) {
	areas[shape] = shape.Area()
	ch <- true
}

func main() {
	shapes := []Shape{
		Square{2},
		Rectangle{2, 3},
		Square{3},
	}

	areas := make(map[Shape]int)
	ch := make(chan bool)

	for _, shape := range shapes {
		go recordArea(shape, areas, ch)
	}

	for range shapes {
		<-ch
	}

	for shape, area := range areas {
		fmt.Println(shape, area)
	}
}
