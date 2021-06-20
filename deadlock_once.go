package main

import (
	"sync"
)

var onceA, onceB sync.Once

func main() {
	var initB func()
	initA := func() { onceB.Do(initB) }
	initB = func() { onceA.Do(initA) }

	onceA.Do(initA)
}

