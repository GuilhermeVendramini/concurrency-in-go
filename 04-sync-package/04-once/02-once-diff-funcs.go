package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int

	fmt.Println("Initial count:", count)

	increment := func() { count++ }
	decrement := func() { count-- }

	var once sync.Once
	once.Do(increment)
	once.Do(decrement)

	fmt.Printf("Count: %d\n", count)
}

/*
	Is it surprising that the output displays 1 and not 0 ? This is because sync.Once only
	counts the number of times Do is called, not how many times unique functions passed
	into Do are called.
*/
