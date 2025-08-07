package main

import (
	"net/url"
	"strings"
)

type errorsMap map[string][]string

func (e errorsMap) Get(field string) string {
	if e == nil {
		e = make(errorsMap)
	}

	errorSlice := e[field]
	if len(errorSlice) == 0 {
		return ""
	}

	return errorSlice[0]
}

func (e errorsMap) Add(field, message string) {
	if e == nil {
		e = make(errorsMap)
	}
	e[field] = append(e[field], message)
}

type Form struct {
	Data      url.Values
	ErrorsMap errorsMap
}

func NewForm(data url.Values) *Form {
	return &Form{
		Data:      data,
		ErrorsMap: errorsMap{},
	}
}

func (f *Form) Has(field string) bool {
	x := f.Data.Get(field)
	return x != ""
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Data.Get(field)
		if len(strings.TrimSpace(value)) == 0 {
			f.ErrorsMap.Add(field, "This field cannot be blank")
		}
	}
}

func (f *Form) Check(ok bool, key, message string) {
	if !ok {
		f.ErrorsMap.Add(key, message)
	}
}

func (f *Form) Valid() bool {
	return len(f.ErrorsMap) == 0
}
