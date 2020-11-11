// go run gotabulate.go
// go get github.com/bndr/gotabulate
package main

import (
	"fmt"

	"github.com/bndr/gotabulate"
)

func main() {
	// Create Some Fake Rows
	row1 := []interface{}{"john", 20, "ready"}
	row2 := []interface{}{"bndr", 23, "ready"}

	// Create an object from 2D interface array
	t := gotabulate.Create([][]interface{}{row1, row2})

	// Set the Headers (optional)
	t.SetHeaders([]string{"age", "status"})

	// Set the Empty String (optional)
	t.SetEmptyString("None")

	// Set Align (Optional)
	t.SetAlign("right")

	// Print the result: grid, or simple
	fmt.Println(t.Render("grid"))
}
