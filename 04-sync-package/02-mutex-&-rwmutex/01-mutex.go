package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()         // <1>
		defer lock.Unlock() // <2>
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	decrement := func() {
		lock.Lock()         // <1>
		defer lock.Unlock() // <2>
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}

	// Increment
	var arithmetic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	// Decrement
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()
	fmt.Println("Arithmetic complete.")
}

/*
	<1> Here we request exclusive use of the critical section—in this case the count vari‐
	able—guarded by a Mutex , lock .

	<2> Here we indicate that we’re done with the critical section lock is guarding.
*/
