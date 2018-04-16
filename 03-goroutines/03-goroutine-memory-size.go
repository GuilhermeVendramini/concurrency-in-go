/*
	In the following example, we combine the fact that goroutines are not garbage collec‐
	ted with the runtime’s ability to introspect upon itself and measure the amount of
	memory allocated before and after goroutine creation.
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() { wg.Done(); <-c } // <1>

	const numGoroutines = 1e4 // <2>
	wg.Add(numGoroutines)
	before := memConsumed() // <3>
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memConsumed() // <4>
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
}

/*
	<1> We require a goroutine that will never exit so that we can keep a number of them
	in memory for measurement. Don’t worry about how we’re achieving this at this
	time; just know that this goroutine won’t exit until the process is finished.

	<2> Here we define the number of goroutines to create. We will use the law of large
	numbers to asymptotically approach the size of a goroutine.

	<3> Here we measure the amount of memory consumed before creating our gorou‐
	tines.

	<4> And here we measure the amount of memory consumed after creating our
	goroutines.
*/
