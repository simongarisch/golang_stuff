package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type DisplacementParameters struct {
	a  float64
	v0 float64
	s0 float64
}

func getUserParameters() (DisplacementParameters, error) {
	var params DisplacementParameters
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter values for acceleration, initial velocity, and initial displacement")
	fmt.Printf("separated by a space: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		return params, err
	}

	if len(input) == 0 {
		return params, errors.New("No text entered")
	}

	inputStrings := strings.Fields(input)
	var inputFloats []float64
	for _, str := range inputStrings {
		floatValue, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return params, err
		}
		inputFloats = append(inputFloats, floatValue)
	}

	if len(inputFloats) != 3 {
		return params, errors.New("Please only enter 3 values")
	}

	params.a = inputFloats[0]
	params.v0 = inputFloats[1]
	params.s0 = inputFloats[2]

	return params, nil
}

func getUserTime() (float64, error) {
	var t float64
	var err error
	fmt.Printf("Please enter a value for time: ")

	var input string
	_, err = fmt.Scan(&input)
	if err != nil {
		return t, err
	}

	return strconv.ParseFloat(input, 64)
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2.0) + v0*t + s0
	}
}

func main() {
	var params DisplacementParameters
	var t float64
	var err error

	params, err = getUserParameters()
	if err != nil {
		panic(err)
	}

	t, err = getUserTime()
	if err != nil {
		panic(err)
	}
	//fmt.Println(params, t)

	fn := GenDisplaceFn(params.a, params.v0, params.s0)
	fmt.Println("The displacement is", fn(t))
}
