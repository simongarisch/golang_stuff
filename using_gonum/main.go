package main

// see the docs https://godoc.org/gonum.org/v1/gonum/mat

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	u := mat.NewVecDense(3, []float64{1, 2, 3})
	println("u: ")
	matPrint(u)

	v := mat.NewVecDense(3, []float64{4, 5, 6})
	println("v: ")
	matPrint(v)

	w := mat.NewVecDense(3, nil)
	w.AddVec(u, v)
	println("u + v: ")
	matPrint(w)

	// Or, you can overwrite u with the result of your operation to
	//save space.
	u.AddVec(u, v)
	println("u (overwritten):")
	matPrint(u)

	// Add u + alpha * v for some scalar alpha
	w.AddScaledVec(u, 2, v)
	println("u + 2 * v: ")
	matPrint(w)

	// Subtract v from u
	w.SubVec(u, v)
	println("v - u: ")
	matPrint(w)

	// Scale u by alpha
	w.ScaleVec(23, u)
	println("u * 23: ")
	matPrint(w)

	// Compute the dot product of u and v
	// Since float64’s don’t have a dot method, this is not done
	//inplace
	d := mat.Dot(u, v)
	println("u dot v: ", d)
	// Find length of v
	l := v.Len()
	println("Length of v: ", l)

	// We can also find the length of v in Euclidean space
	// The 2 parameter specifices that this is the Frobenius norm
	// Rather than the maximum absolute column sum
	// This distinction is more important when Norm is applied to
	// Matrices since vectors only have one column and the the
	// maximum absolute column sum is the Frobenius norm squared.
	println(mat.Norm(v, 2))
}

func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}
