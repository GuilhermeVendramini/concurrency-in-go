package main

import (
	"bytes"
	"fmt"
	"sync"
)

func main() {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()

		var buff bytes.Buffer
		for i, b := range data {
			fmt.Println("Item:", i)
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")
	go printData(&wg, data[:3]) // <1>
	go printData(&wg, data[3:]) // <2>

	wg.Wait()
}

/*
	<1> Here we pass in a slice containing the first three bytes in the data structure.

	<2> Here we pass in a slice containing the last three bytes in the data structure.

	In this example, you can see that because printData doesn’t close around the data
	slice, it cannot access it, and needs to take in a slice of byte to operate on. We pass in
	different subsets of the slice, thus constraining the goroutines we start to only the part
	of the slice we’re passing in. Because of the lexical scope, we’ve made it impossible to
	do the wrong thing, and so we don’t need to synchronize memory access or share
	data through communication.
*/
