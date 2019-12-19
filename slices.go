package main

import "fmt"

// arrays have a fixed size
// but we can change the size of slices

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

	a1 := [3]string{"a", "b", "c"}
	sli1 := a1[0:1]
	fmt.Println(len(sli1), cap(sli1)) // 1, 3 -> length and capacity
	fmt.Println("----")

	// variable slices
	sli := make([]int, 10)
	fmt.Println(sli)                // [0 0 0 0 0 0 0 0 0 0]
	fmt.Println(len(sli), cap(sli)) // 10, 10

	// there is a time penalty if you append to a slice beyond allocated capacity
	// as a new slice has to be created and returned
}
