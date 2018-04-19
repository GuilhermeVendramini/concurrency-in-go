package main

import (
	"fmt"
	"net/http"
)

func main() {
	type Result struct { // <1>
		Error    error
		Response *http.Response
	}
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result { // <2>
		results := make(chan Result)
		go func() {
			defer close(results)

			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				result = Result{Error: err, Response: resp} // <3>
				select {
				case <-done:
					return
				case results <- result: // <4>
				}
			}
		}()
		return results
	}

	done := make(chan interface{})
	defer close(done)

	urls := []string{"https://www.google.com", "https://badhost", "https://www.bing.com"}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil { // <5>
			fmt.Printf("error: %v\n", result.Error)
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}

/*
	<1> Here we create a type that encompasses both the *http.Response and the error
	possible from an iteration of the loop within our goroutine.

	<2> This line returns a channel that can be read from to retrieve results of an iteration
	of our loop.

	<3> Here we create a Result instance with the Error and Response fields set.

	<4> This is where we write the Result to our channel.

	<5> Here, in our main goroutine, we are able to deal with errors coming out of the
	goroutine started by checkStatus intelligently, and with the full context of the
	larger program.

	The key thing to note here is how we’ve coupled the potential result with the potential
	error. This represents the complete set of possible outcomes created from the gorou‐
	tine checkStatus , and allows our main goroutine to make decisions about what to do
	when errors occur. In broader terms, we’ve successfully separated the concerns of
	error handling from our producer goroutine. This is desirable because the goroutine
	that spawned the producer goroutine—in this case our main goroutine—has more
	context about the running program, and can make more intelligent decisions about
	what to do with errors.
*/
