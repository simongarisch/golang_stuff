package main

import (
	"os"
	"testing"
)

var app application

func TestMain(m *testing.M) {
	// always be executed before the tests run
	os.Exit(m.Run())
}
