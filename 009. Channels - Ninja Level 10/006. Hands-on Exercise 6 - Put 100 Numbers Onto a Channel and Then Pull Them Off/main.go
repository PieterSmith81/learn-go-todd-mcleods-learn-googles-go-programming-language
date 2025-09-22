package main

import "fmt"

func main() {
	c := make(chan int)

	// Send
	go func() {
		for i := 0; i < 100; i++ { // Or you can use: for i := range 100
			c <- i
		}

		close(c)
	}()

	// Receive
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("About to exit.")
}
