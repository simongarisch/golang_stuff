package main

// go run example.go
// go get https://github.com/juju/errors

import (
	"fmt"

	"github.com/juju/errors"
)

var err = errors.New("some error message")

func getError() error {
	return err
}

func getErrorTrace() error {
	return errors.Trace(err)
}

func main() {
	err1 := getError()
	err2 := getErrorTrace()
	fmt.Println(err1) // some error message
	fmt.Println("********")
	fmt.Println(err2) // some error message
	if err2 != nil {
		fmt.Println(errors.ErrorStack(err))
		// .../example.go:12: some error message
	}

	fmt.Println("*****************")
	err := errors.Errorf("original")
	err = errors.Annotatef(err, "context")
	err = errors.Annotatef(err, "more context")
	fmt.Println(err.Error()) // "more context: context: original"
}
