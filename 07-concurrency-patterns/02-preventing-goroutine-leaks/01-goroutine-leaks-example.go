package main

import "fmt"

func main() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for s := range strings {
				// Do something interesting
				fmt.Println(s)
			}
		}()
		return completed
	}

	doWork(nil)
	// Perhaps more work is done here
	fmt.Println("Done.")
}

/*
	Here we see that the main goroutine passes a nil channel into doWork . Therefore, the
	strings channel will never actually gets any strings written onto it, and the goroutine
	containing doWork will remain in memory for the lifetime of this process (we would
	even deadlock if we joined the goroutine within doWork and the main goroutine).
	In this example, the lifetime of the process is very short, but in a real program, gorou‐
	tines could easily be started at the beginning of a long-lived program. In the worst
	case, the main goroutine could continue to spin up goroutines throughout its life,
	causing creep in memory utilization.
	The way to successfully mitigate this is to establish a signal between the parent gorou‐
	tine and its children that allows the parent to signal cancellation to its children. By
	convention, this signal is usually a read-only channel named done . The parent gorou‐
	tine passes this channel to the child goroutine and then closes the channel when it
	wants to cancel the child goroutine. See the example "02-goroutine-leaks-cancellation.go".
*/
