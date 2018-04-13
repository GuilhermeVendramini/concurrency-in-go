package main

import (
	"fmt"
	"sync"
)

func main() {
	var memoryAccess sync.Mutex // <1>
	var value int
	go func() {
		memoryAccess.Lock() // <2>
		value++
		memoryAccess.Unlock() // <3>
	}()

	memoryAccess.Lock() // <4>
	if value == 0 {
		fmt.Printf("the value is %v.\n", value)
	} else {
		fmt.Printf("the value is %v.\n", value)
	}
	memoryAccess.Unlock() // <5>
}

/*
<1> Here we add a variable that will allow our code to synchronize access to the data
variable’s memory.

<2> Here we declare that until we declare otherwise, our goroutine should have
exclusive access to this memory.

<3> Here we declare that the goroutine is done with this memory.

<4> Here we once again declare that the following conditional statements should have
exclusive access to the data variable’s memory.

<5> Here we declare we’re once again done with this memory.

You may have noticed that while we have solved our data race, we haven’t actually
solved our race condition! The order of operations in this program is still nondeter‐
ministic; we’ve just narrowed the scope of the nondeterminism a bit. In this example,
either the goroutine will execute first, or both our if and else blocks will. We still
don’t know which will occur first in any given execution of this program. Later, we’ll
explore the tools to solve this kind of issue properly.

On its face this seems pretty simple: if you find you have critical sections, add points
to synchronize access to the memory! Easy, right? Well...sort of.

It is true that you can solve some problems by synchronizing access to the memory,
but as we just saw, it doesn’t automatically solve data races or logical correctness. Fur‐
ther, it can also create maintenance and performance problems.

By Katherine Cox-Buday
*/
