package main

import (
	"fmt"
)

func main() {
	data := make([]int, 4)

	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			handleData <- data[i]
		}
	}

	handleData := make(chan int)
	go loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}
}

/*
	We can see that the data slice of integers is available from both the loopData function
	and the loop over the handleData channel; however, by convention we’re only access‐
	ing it from the loopData function. But as the code is touched by many people, and
	deadlines loom, mistakes might be made, and the confinement might break down
	and cause issues. As I mentioned, a static-analysis tool might catch these kinds of
	issues, but static analysis on a Go codebase suggests a level of maturity that not many
	teams achieve. This is why I prefer lexical confinement: it wields the compiler to
	enforce the confinement.
*/
