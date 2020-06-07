// The bridge pattern is about preferring composition over inheritance
package main

import (
	"fmt"
	"reflect"
)

type Colour struct {
	Colour string
}

type Shape interface {
	Draw()
}

type Rectangle struct{ colour Colour }
type Square struct{ colour Colour }

func (r *Rectangle) Draw() {
	fmt.Println(reflect.TypeOf(*r), r.colour)
}

func (s *Square) Draw() {
	fmt.Println(reflect.TypeOf(*s), s.colour)
}

func main() {
	red := Colour{"Red"}
	green := Colour{"Green"}
	r := Rectangle{red}
	s := Square{green}
	r.Draw() // main.Rectangle {Red}
	s.Draw() // main.Square {Green}
}
