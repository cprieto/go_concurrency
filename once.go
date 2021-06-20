package main

import (
	"fmt"
	"sync"
)

var count int
var wg sync.WaitGroup
var once sync.Once

func main() {
	increment := func() {
		count++
	}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			once.Do(increment)
		}()
	}

	wg.Wait()
	fmt.Printf("count: %v\n", count)

}
