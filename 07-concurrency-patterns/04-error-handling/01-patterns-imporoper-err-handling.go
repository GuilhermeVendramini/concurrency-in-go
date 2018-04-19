package main

import (
	"fmt"
	"net/http"
)

func main() {
	checkStatus := func(
		done <-chan interface{},
		urls ...string,
	) <-chan *http.Response {
		responses := make(chan *http.Response)
		go func() {
			defer close(responses)
			for _, url := range urls {
				resp, err := http.Get(url)
				if err != nil {
					fmt.Println(err) // <1>
					continue
				}
				select {
				case <-done:
					return
				case responses <- resp:
				}
			}
		}()
		return responses
	}

	done := make(chan interface{})
	defer close(done)

	urls := []string{"https://www.google.com", "https://badhost"}
	for response := range checkStatus(done, urls...) {
		fmt.Printf("Response: %v\n", response.Status)
	}
}

/*
	<1> Here we see the goroutine doing its best to signal that there’s an error. What else
	can it do? It can’t pass it back! How many errors is too many? Does it continue
	making requests?

	Here we see that the goroutine has been given no choice in the matter. It can’t simply
	swallow the error, and so it does the only sensible thing: it prints the error and hopes
	something is paying attention. Don’t put your goroutines in this awkward position. I
	suggest you separate your concerns: in general, your concurrent processes should
	send their errors to another part of your program that has complete information
	about the state of your program, and can make a more informed decision about what
	to do. See the correct solution to this problem at 02-patterns-proper-err-handling.go.
*/
