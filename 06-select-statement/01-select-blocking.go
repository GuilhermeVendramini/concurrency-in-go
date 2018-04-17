package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(c) // <1>
	}()

	fmt.Println("Blocking on read...")
	select {
	case <-c: // <2>
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	}
}

/*
	<1> Here we close the channel after waiting five seconds.

	<2> Here we attempt a read on the channel. Note that as this code is written, we don’t
	require a select statement—we could simply write <-c —but we’ll expand on this
	example.
*/
