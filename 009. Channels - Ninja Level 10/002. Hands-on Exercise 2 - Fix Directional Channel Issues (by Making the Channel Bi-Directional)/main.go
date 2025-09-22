package main

import (
	"fmt"
)

func main() {
	// cs := make(chan<- int) // Original broken send channel directional code (problem 1).
	// cr := make(<-chan int) // Original broken receive channel directional code (problem 2).

	cs := make(chan int) // Both these problems can be fixed by making the channel bi-directional.

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
