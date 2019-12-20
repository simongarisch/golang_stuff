package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Question struct {
	Subject string
	Text    string
}

func main() {

	questions := []Question{
		Question{
			Subject: "name",
			Text:    "Please enter your first name...",
		},
		Question{
			Subject: "address",
			Text:    "Please enter your address...",
		},
	}
	//q1 := questions[0]
	//fmt.Println(q1, q1.Subject, q1.Text)

	// get the user details
	user := make(map[string]string)
	reader := bufio.NewReader(os.Stdin)

	for _, question := range questions {
		fmt.Println(question.Text)
		input, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		input = strings.TrimRight(input, "\r\n")
		user[question.Subject] = input
	}
	//fmt.Println(user)

	// crate a json object from the map
	byteArr, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	//fmt.Println(byteArr) // [123 34 97 ... a byte array

	// now print this back to
	newUser := make(map[string]string)
	err = json.Unmarshal(byteArr, &newUser)
	if err != nil {
		panic(err)
	}
	fmt.Println(newUser)
}
