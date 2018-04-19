package main

import (
	"fmt"
	"math/rand"
)

func main() {
	newRandStream := func() <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.") // <1>
			defer close(randStream)
			for {
				randStream <- rand.Int()
			}
		}()

		return randStream
	}

	randStream := newRandStream()
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
}

/*
	<1> Here we print out a message when the goroutine successfully terminates.

	You can see from the output that the deferred fmt.Println statement never gets run.
	After the third iteration of our loop, our goroutine blocks trying to send the next ran‐
	dom integer to a channel that is no longer being read from. We have no way of telling
	the producer it can stop. The solution, just like for the receiving case, is to provide the
	producer goroutine with a channel informing it to exit. See example "04-leak-from-blocked-channel-write-solved.go".
*/
