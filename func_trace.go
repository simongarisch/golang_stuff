// go run func_trace.go
package main

import (
	"fmt"
	"runtime"
)

func trace() (string, int, string) {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return "?", 0, "?"
	}

	fn := runtime.FuncForPC(pc)
	return file, line, fn.Name()
}

func testTrace() {
	file, line, fnname := trace()

	fmt.Println("file: ", file)     // file:  C:/Users/.../func_trace.go
	fmt.Println("line: ", line)     // line:  20
	fmt.Println("fnname: ", fnname) // fnname:  main.testTrace
}

func main() {
	testTrace()
}
