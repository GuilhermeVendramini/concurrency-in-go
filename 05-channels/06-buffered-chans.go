package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var stdoutBuff bytes.Buffer         // <1>
	defer stdoutBuff.WriteTo(os.Stdout) // <2>

	intStream := make(chan int, 4) // <3>
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done.")
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received %v.\n", integer)
	}
}

/*
	<1> Here we create an in-memory buffer to help mitigate the nondeterministic nature
	of the output. It doesn’t give us any guarantees, but it’s a little faster than writing
	to stdout directly.

	<2> Here we ensure that the buffer is written out to stdout before the process exits.

	<3> Here we create a buffered channel with a capacity of one.
*/
