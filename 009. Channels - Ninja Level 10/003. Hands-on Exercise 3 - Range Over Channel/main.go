package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)

	// Note the use of the go keyword along with a
	// self executing anonymous function (a.k.a. function literal) below.
	// This combination sends this code off into its own go routine which stops it from blocking,
	// and allows the receive() function to run and start pulling values off the channel.
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}

		close(c)
	}()

	return c
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
