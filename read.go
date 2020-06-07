package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func IsEven(number int) bool {
	return number%2 == 0
}

func ReadNamesFile(fileName string) ([]string, error) {
	bytesArr, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	fileSlice := strings.Fields(string(bytesArr))
	numNames := int(len(fileSlice) / 2)
	if !IsEven(numNames) { // entries must be even for fname, lname
		return nil, errors.New("File format must be 'fname lname'")
	}

	return fileSlice, nil
}

func main() {
	// prompt the user for a file name
	var fileName string
	//fileName = "names.txt"

	fmt.Printf("Please enter a file name to read: ")
	_, err := fmt.Scan(&fileName)
	if err != nil {
		panic(err)
	}

	// read the file
	fileSlice, err := ReadNamesFile(fileName)
	if err != nil {
		panic(err)
	}
	//fmt.Println(fileSlice)

	// create a slice of Name structs
	var names []Name
	numNames := int(len(fileSlice) / 2)
	var fname, lname string
	for i := 0; i < numNames; i++ {
		fname = fileSlice[i*2]
		lname = fileSlice[i*2+1]
		//fmt.Println(fname, lname)
		names = append(names, Name{fname, lname})
	}
	//fmt.Println(names)

	// print these out
	for _, name := range names {
		fmt.Println(name.fname, name.lname)
	}

}
