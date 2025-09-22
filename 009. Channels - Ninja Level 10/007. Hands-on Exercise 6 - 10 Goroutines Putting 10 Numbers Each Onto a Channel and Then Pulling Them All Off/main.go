package main

import "fmt"

func main() {
	c := make(chan int)

	// Send
	for i := 0; i < 10; i++ { // Or you can use: for range 10
		go func() {
			for i := 0; i < 10; i++ { // Or you can use: for i := range 10
				c <- i
			}
		}()
	}

	// Receive
	for i := 0; i < 100; i++ { // Or you can use: for i := range 100
		fmt.Println(i, "\t", <-c)
	}

	fmt.Println("About to exit.")
}
