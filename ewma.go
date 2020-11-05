// go run ewma.go
package main

import (
	"fmt"

	"github.com/VividCortex/ewma"
)

func main() {
	samples := [100]float64{
		4599, 5711, 4746, 4621, 5037, 4218, 4925, 4281, 5207, 5203, 5594, 5149,
	}

	e := ewma.NewMovingAverage()  //=> Returns a SimpleEWMA if called without params
	a := ewma.NewMovingAverage(5) //=> returns a VariableEWMA with a decay of 2 / (5 + 1)

	for _, f := range samples {
		e.Add(f)
		a.Add(f)
	}

	fmt.Println(e.Value()) //=> 13.577404704631077
	fmt.Println(a.Value()) //=> 1.5806140565521463e-12
}
