package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1) // <1>
	go func() {
		defer wg.Done() // <2>
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1)
	}()

	wg.Add(1) // <1>
	go func() {
		defer wg.Done() // <2>
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(2)
	}()

	wg.Wait() // <3>
	fmt.Println("All goroutines complete.")
}

/*
	<1> Here we call Add with an argument of 1 to indicate that one goroutine is begin‐
	ning.

	<2> Here we call Done using the defer keyword to ensure that before we exit the
	goroutine’s closure, we indicate to the WaitGroup that we’ve exited.

	<3> Here we call Wait , which will block the main goroutine until all goroutines have
	indicated they have exited.
*/
