package main

import (
	"fmt"

	"github.com/jinzhu/copier"
)

type MyStruct struct {
	Data string
}

func (orig MyStruct) CreateMemento() func(*MyStruct) {
	pmemento := &MyStruct{}
	copier.Copy(pmemento, &orig)
	return func(m *MyStruct) {
		copier.Copy(m, pmemento)
	}
}

func main() {
	s := MyStruct{"with some string data"}
	fmt.Println(s) // {with some string data}

	restore := s.CreateMemento()
	s.Data = "some new data"
	fmt.Println(s) // {some new data}
	restore(&s)
	fmt.Println(s) // {with some string data}
}
