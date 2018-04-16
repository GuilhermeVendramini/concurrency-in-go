package main

import (
	"fmt"
	"sync"
)

func main() {
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin // <1>
			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutines...")
	close(begin) // <2>
	wg.Wait()
}

/*
	<1> Here the goroutine waits until it is told it can continue.
	<2> Here we close the channel, thus unblocking all the goroutines simultaneously.
*/
