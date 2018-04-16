package main

import (
	"fmt"
)

func main() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello channels!" // <1>
	}()
	fmt.Println(<-stringStream) // <2>
}

/*
	<1> Here we pass a string literal onto the channel stringStream .
	<2> Here we read the string literal off of the channel and print it out to stdout .
*/
