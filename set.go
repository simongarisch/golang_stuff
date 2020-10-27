// go run set.go
// https://github.com/Workiva/go-datastructures/blob/master/set/dict_test.go
package main

import (
	"fmt"

	"github.com/Workiva/go-datastructures/set"
)

func main() {
	s := set.New()
	s.Add("this")
	s.Add("and")
	s.Add("this")
	s.Add(1)
	s.Add(1)

	fmt.Println(s.Len())         // 3
	fmt.Println(s.Exists("and")) // true
	fmt.Println(s.Exists("xxx")) // false
	fmt.Println(s)               // &{map[1:{} and:{} this:{}] {{0 0} 0 0 0 0} []}
	s.Remove("and")
	fmt.Println(s) // &{map[1:{} this:{}] {{0 0} 0 0 0 0} []}
	s.Clear()
	fmt.Println(s) // &{map[] {{0 0} 0 0 0 0} []}
}
