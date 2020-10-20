// go run reflect_clean_struct.go
package main

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

func CleanString(s string) string {
	return strings.ToUpper(strings.TrimSpace(s))
}

func CleanStruct(data interface{}) {
	values := reflect.ValueOf(data)
	for i := 0; i < reflect.Indirect(values).NumField(); i++ {
		v := reflect.Indirect(values).Field(i)
		if v.Kind() == reflect.String {
			newValue := CleanString(v.String())
			v.SetString(newValue)
		}
		if v.Type().String() == "sql.NullString" {
			ns := v.Interface().(sql.NullString)
			if ns.Valid {
				ns.String = CleanString(ns.String)
				v.Set(reflect.ValueOf(ns))
			}
		}
	}
}

type TestStruct struct {
	A string
	B int
	C sql.NullString
	D string
	E float64
}

func main() {
	s := TestStruct{"this", 7, sql.NullString{"some_value", true}, "is", 7.77}
	CleanStruct(&s)
	fmt.Println(s) // {THIS 7 {SOME_VALUE true} IS 7.77}
}
