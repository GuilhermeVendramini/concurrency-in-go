package main

import (
	"fmt"
)

func main() {
	intStream := make(chan int)
	go func() {
		defer close(intStream) // <1>
		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()

	for integer := range intStream { // <2>
		fmt.Printf("%v ", integer)
	}
}

/*
	<1> Here we ensure that the channel is closed before we exit the goroutine. This is a
	very common pattern.

	<2> Here we range over intStream.
*/
