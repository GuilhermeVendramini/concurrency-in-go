package main

import (
	"fmt"
	"sync"
)

func main() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}

	myPool.Get()             // <1>
	instance := myPool.Get() // <1>
	myPool.Put(instance)     // <2>
	myPool.Get()             // <3>
}

/*
	<1> Here we call Get on the pool. These calls will invoke the New function defined on
	the pool since instances haven’t yet been instantiated.

	<2> Here we put an instance previously retrieved back in the pool. This increases the
	available number of instances to one.

	<3> When this call is executed, we will reuse the instance previously allocated and put
	it back in the pool. The New function will not be invoked.
*/
