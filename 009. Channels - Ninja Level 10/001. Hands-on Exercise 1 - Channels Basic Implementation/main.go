package main

import (
	"fmt"
)

func main() {
	// Original, blocking code.
	// c := make(chan int)
	// c <- 42 // Blocks here, since the is no channel receiver to simultaneously take (a.k.a. receive) the value off the channel.
	// fmt.Println(<-c) // So, the code will never reach this channel receiver.

	// Channels basic implementation using a function literal
	// (a.k.a. anonymous, self-executing function).
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

	// Channels basic implementation using a buffered channel.
	c2 := make(chan int, 1)

	c2 <- 42

	fmt.Println(<-c2)
}
