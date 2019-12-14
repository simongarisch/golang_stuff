package main

import "fmt"

func main() {
	// a slice is a sequence of items
	loons := []string{"bugs", "daffy", "taz"}
	fmt.Println("loons = %v (type %T)\n", loons, loons)

	fmt.Println(len(loons))
	fmt.Println("----")

	fmt.Println(loons[1])
	fmt.Println("----")

	fmt.Println(loons[1:])
	fmt.Println("----")

	for index, name := range loons {
		fmt.Printf("%s at %d\n", name, index)
	}
	fmt.Println("----")

	loons = append(loons, "elmer")
	for index, name := range loons {
		fmt.Printf("%s at %d\n", name, index)
	}
	fmt.Println("----")
}
