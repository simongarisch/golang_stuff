package main

import "fmt"

type Square struct {
	Width int
}

func (s Square) Area() int {
	return s.Width * s.Width
}

func recordArea(s Square, areas map[Square]int, c chan bool) {
	areas[s] = s.Area()
	c <- true
}

func main() {
	squares := []Square{
		Square{1},
		Square{2},
		Square{3},
	}

	areas := make(map[Square]int)
	c := make(chan bool)

	for _, s := range squares {
		go recordArea(s, areas, c)
	}

	for i := 0; i < len(squares); i++ {
		<-c
	}

	fmt.Println(areas)
}
