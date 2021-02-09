// go run util.go
// there look to be a number of github projects devoted to go util functions.
// https://github.com/subchen/go-stack
// https://github.com/gxxgle/go-utils
// https://github.com/zillow/godash
// https://github.com/go-goodies/go_utils
package main

import (
	"database/sql"
	"fmt"
	"math"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"time"
)

// Round2dp rounds some input float to 2dp.
func Round2dp(x float64) float64 {
	return math.Round(x*100) / 100
}

// SliceContains returns true if an item exists in a slice, false otherwise.
func SliceContains(slice []interface{}, elem interface{}) bool {
	for _, item := range slice {
		if item == elem {
			return true
		}
	}
	return false
}

// StringSliceContains returns true if an item exists in a string slice, false otherwise.
func StringSliceContains(values []string, value string) bool {
	for _, item := range values {
		if item == value {
			return true
		}
	}
	return false
}

// CleanString cleans and returns some input string.
func CleanString(s string) string {
	return strings.ToUpper(strings.TrimSpace(s))
}

// SetFromStringSlice returns only unique values from some slice.
func SetFromStringSlice(slice []string) []string {
	var result []string
	contains := make(map[string]bool)
	for _, item := range slice {
		item = CleanString(item)
		_, ok := contains[item]
		if !ok {
			result = append(result, item)
		}
		contains[item] = true
	}
	sort.Strings(result)
	return result
}

// CleanStruct cleans all string fields in a struct
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

// GetSydneyTime gets the time in Sydney, Australia.
func GetSydneyTime() (time.Time, error) {
	currentTime := time.Now()
	location, err := time.LoadLocation("Australia/Sydney")
	if err != nil {
		return currentTime, err
	}

	sydneyTime := currentTime.In(location)
	return sydneyTime, nil
}

// DatesEqual returns true if two dates are equal, false otherwise
func DatesEqual(date1 time.Time, date2 time.Time) bool {
	year1, month1, day1 := date1.Date()
	year2, month2, day2 := date2.Date()
	if year1 == year2 {
		if month1 == month2 {
			if day1 == day2 {
				return true
			}
		}
	}
	return false
}

// IsWeekend returns true if some date falls on a weekend, false otherwise.
func IsWeekend(date time.Time) bool {
	weekday := date.Weekday()
	if weekday == time.Saturday || weekday == time.Sunday {
		return true
	}
	return false
}

// GetWorkingDaysUntil returns the number of working days
// until a particular date.
func GetWorkingDaysUntil(date time.Time) (int, error) {
	var days int
	currentTime, err := GetSydneyTime()
	if err != nil {
		return days, err
	}

	if date.Before(currentTime) {
		return days, nil
	}
	if DatesEqual(date, currentTime) {
		return days, nil
	}

	// keep on counting while the day is not a weekend...
	for {
		if !IsWeekend(currentTime) {
			days++
		}
		currentTime = currentTime.Add(time.Hour * 24)
		if DatesEqual(date, currentTime) {
			break
		}
	}
	return days, nil
}

// Trace provides some function traceback information.
func Trace() (string, int, string) {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return "?", 0, "?"
	}

	fn := runtime.FuncForPC(pc)
	return file, line, fn.Name()
}

// PrintStringFloatMap displays a map for debugging.
func PrintStringFloatMap(input map[string]float64) {
	var keys []string
	for key := range input {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		value := fmt.Sprintf("%f", input[key])
		fmt.Println(key + ": " + value)
	}
}

// RemoveStringFromSlice takes a slice of strings and removes all
// instances of a specific value.
func RemoveStringFromSlice(toRemove string, slice []string) []string {
	var newSlice []string
	for _, value := range slice {
		if value != toRemove {
			newSlice = append(newSlice, value)
		}
	}
	return newSlice
}

func main() {
	fmt.Println(Round2dp(2.3333))                                  // 2.33
	fmt.Println(SliceContains([]interface{}{"a", 2, "c"}, "c"))    // true
	fmt.Println(StringSliceContains([]string{"a", "b", "c"}, "c")) // true
	fmt.Println(CleanString(" xxx "))                              // 'XXX'
	fmt.Println(SetFromStringSlice([]string{"a", "a", "b"}))       // [A B]

	type TestStruct struct {
		A string
		B int
		C string
		D sql.NullString
	}

	s := TestStruct{"one ", 2, " three", sql.NullString{"some value", true}}
	CleanStruct(&s)
	fmt.Println(s) // {ONE 2 THREE {SOME VALUE true}}

	currentTime, err := GetSydneyTime()
	if err != nil {
		panic(err)
	}

	fmt.Println(currentTime)                          // 2020-12-09 08:31:22.8109344 +1100 AEDT
	fmt.Println(DatesEqual(currentTime, currentTime)) // true
	fmt.Println(IsWeekend(currentTime))               // false

	currentTime = currentTime.Add(time.Hour * 24 * 7) // 7 days, 5 working days
	fmt.Println(GetWorkingDaysUntil(currentTime))     // 5
}
