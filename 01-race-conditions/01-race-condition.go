package main

import (
	"fmt"
)

func main() {
	var data int
	go func() {
		data++
	}()
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	}
}

/*
	Here, lines 10 and 12 are both trying to access the variable data, but there is no guaran‐
	tee what order this might happen in. There are three possible outcomes to running
	this code:
	• Nothing is printed. In this case, line 10 was executed before line 12.
	• “the value is 0” is printed. In this case, lines 12 and 13 were executed before line 10.
	• “the value is 1” is printed. In this case, line 12 was executed before line 10, but line 10
	was executed before line 13.
*/
