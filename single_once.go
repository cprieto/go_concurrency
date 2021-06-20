package main

import (
	"fmt"
	"sync"
)

var count int
var once sync.Once

func main() {
	increment := func() {
		count++
	}

	decrement := func() {
		count--
	}

	once.Do(increment)
	once.Do(decrement)

	// Once works per instance not per func!
	fmt.Printf("count: %v\n", count)
}
