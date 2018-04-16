package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})    // <1>
	queue := make([]interface{}, 0, 10) // <2>

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()        // <8>
		queue = queue[1:] // <9>
		fmt.Println("Removed from queue")
		c.L.Unlock() // <10>
		c.Signal()   // <11>
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()            // <3>
		for len(queue) == 2 { // <4>
			c.Wait() // <5>
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second) // <6>
		c.L.Unlock()                        // <7>
	}
}

/*
	<1> First, we create our condition using a standard sync.Mutex as the Locker .

	<2> Next, we create a slice with a length of zero. Since we know we’ll eventually add
	10 items, we instantiate it with a capacity of 10.

	<3> We enter the critical section for the condition by calling Lock on the condition’s
	Locker .

	<4> Here we check the length of the queue in a loop. This is important because a sig‐
	nal on the condition doesn’t necessarily mean what you’ve been waiting for has
	occurred—only that something has occurred.

	<5> We call Wait , which will suspend the main goroutine until a signal on the condi‐
	tion has been sent.

	<6> Here we create a new goroutine that will dequeue an element after one second.

	<7> Here we exit the condition’s critical section since we’ve successfully enqueued an
	item.

	<8> We once again enter the critical section for the condition so we can modify data
	pertinent to the condition.

	<9> Here we simulate dequeuing an item by reassigning the head of the slice to the
	second item.

	<10> Here we exit the condition’s critical section since we’ve successfully dequeued an
	item.

	<11> Here we let a goroutine waiting on the condition know that something has
	occurred.
*/
