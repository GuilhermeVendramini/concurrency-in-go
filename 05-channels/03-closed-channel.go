package main

import (
	"fmt"
)

func main() {
	intStream := make(chan int)
	close(intStream)
	integer, ok := <-intStream // <1>
	fmt.Printf("(%v): %v", ok, integer)
}

/*
	<1> Here we read from a closed stream.
*/
