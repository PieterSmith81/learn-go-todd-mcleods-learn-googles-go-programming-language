package main

import (
	"fmt"
	"math/rand"
)

func main() {
	c := 0

	for i := 0; i < 100; i++ {
		fmt.Println("Iteration number", i)

		// "statement statement" idiom example
		if x := rand.Intn(5); x == 3 {
			fmt.Println("x is 3.")
			c++
		}
	}

	fmt.Printf("x was three %v times in one hundred loops.", c)
}
