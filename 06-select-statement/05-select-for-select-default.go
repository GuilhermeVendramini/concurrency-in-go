package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	workCounter := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}

		// Simulate work
		workCounter++
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workCounter)
}

/*
	This produces:
	Achieved 5 cycles of work before signalled to stop.
	In this case, we have a loop that is doing some kind of work and occasionally check‐
	ing whether it should stop.
	Finally, there is a special case for empty select statements: select statements with
	no case clauses. These look like this:
	select {}
	This statement will simply block forever.
*/
