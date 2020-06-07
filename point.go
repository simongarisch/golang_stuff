package main

import (
	"fmt"
	"os"
)

type Point struct {
	X int
	Y int
}

func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

type Square struct {
	Center Point
	Length int
}

func NewSquare(center Point, length int) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("length must be > 0 (was %d)", length)
	}

	square := &Square{
		Center: center,
		Length: length,
	}
	return square, nil
}

func (s *Square) Move(dx int, dy int) {
	s.Center.Move(dx, dy)
}

func (s *Square) Area() int {
	return s.Length * s.Length
}

func main() {
	point := Point{0, 0}
	square, err := NewSquare(point, 10)
	if err != nil {
		fmt.Printf("error: can't create square - %s\n", err)
		os.Exit(1)
	}
	fmt.Println(square)

	square.Move(5, 5)
	fmt.Println(square)

	fmt.Println(square.Area())
}
