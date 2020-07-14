/*
You may have noticed that while we have solved our data race, we haven’t actually solved our race condition!
The order of operations in this program is still nondeterministic;
we’ve just narrowed the scope of the nondeterminism a bit.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	var memoryAccess sync.Mutex
	var value int
	go func() {
		memoryAccess.Lock()
		value++
		memoryAccess.Unlock()
	}()

	memoryAccess.Lock()
	fmt.Printf("the value is %v.\n", value)
	memoryAccess.Unlock()
	// the value is 0.
}
