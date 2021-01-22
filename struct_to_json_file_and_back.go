// go run struct_to_json_file_and_back.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Salary struct {
	Basic, HRA, TA float64
}

type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
}

func main() {

	// create our data
	// https://www.golangprograms.com/golang-writing-struct-to-json-file.html
	employee := Employee{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []Salary{
			Salary{
				Basic: 15000.00,
				HRA:   5000.00,
				TA:    2000.00,
			},
			Salary{
				Basic: 16000.00,
				HRA:   5000.00,
				TA:    2100.00,
			},
			Salary{
				Basic: 17000.00,
				HRA:   5000.00,
				TA:    2200.00,
			},
		},
	}

	// write this data to a json file...
	// https://golang.org/pkg/encoding/json/#MarshalIndent
	file, err := json.MarshalIndent(employee, "", " ")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("etest.json", file, 0644)
	if err != nil {
		panic(err)
	}

	// and read this data back
	filename, err := os.Open("etest.json")
	if err != nil {
		log.Fatal(err)
	}
	defer filename.Close()

	data, err := ioutil.ReadAll(filename)
	if err != nil {
		panic(err)
	}

	var result Employee
	err = json.Unmarshal(data, &result)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
