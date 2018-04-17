package main

import (
	"fmt"
)

func main() {
	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5) // <1>
		go func() {                       // <2>
			defer close(resultStream) // <3>
			for i := 0; i <= 5; i++ {
				resultStream <- i
			}
		}()
		return resultStream // <4>
	}

	resultStream := chanOwner()
	for result := range resultStream { // <5>
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving!")
}

/*
	<1> Here we instantiate a buffered channel. Since we know we’ll produce six results,
	we create a buffered channel of five so that the goroutine can complete as quickly
	as possible.

	<2> Here we start an anonymous goroutine that performs writes on resultStream .
	Notice that we’ve inverted how we create goroutines. It is now encapsulated
	within the surrounding function.

	<3> Here we ensure resultStream is closed once we’re finished with it. As the chan‐
	nel owner, this is our responsibility.

	<4> Here we return the channel. Since the return value is declared as a read-only
	channel, resultStream will implicitly be converted to read-only for consumers.

	<5> Here we range over resultStream . As a consumer, we are only concerned with
	blocking and closed channels.
*/
