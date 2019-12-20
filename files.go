package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

func main() {
	// no explicit open / close: ReadFile does this for you
	data, err := ioutil.ReadFile("text.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(data) // [116 104 105 115 3 ... a byte array

	dataString := string(data)
	fmt.Println(dataString)
	fmt.Println("--------")

	// for a .csv file
	data, err = ioutil.ReadFile("test.csv")
	if err != nil {
		panic(err)
	}

	dataString = string(data)
	fmt.Println(dataString)
	fmt.Println("--------")

	// or we can read this csv file with csv.Read
	filePath := "test.csv"
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	fmt.Println(records)                 // [[1 2 3 4] [5 6 7 8] [1 2 3 4]]
	fmt.Println(reflect.TypeOf(records)) // [][]string
	fmt.Println(records[1][1])           // 6

	// writing to a file
	sdat := "Hello, world!"
	err = ioutil.WriteFile("outfile.txt", []byte(sdat), 0777)
	if err != nil {
		panic(err)
	}
}
