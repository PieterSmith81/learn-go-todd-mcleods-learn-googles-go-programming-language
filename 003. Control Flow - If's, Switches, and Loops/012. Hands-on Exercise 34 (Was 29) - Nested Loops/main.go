package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("Outer loop iteration: %v\n", i)

		for j := 0; j < 5; j++ {
			fmt.Printf("\tInner loop iteration: %v\n", j)
		}
	}
}
