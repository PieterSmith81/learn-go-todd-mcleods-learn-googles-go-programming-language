package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i < 42; i++ {
		fmt.Printf("Iteration number %v:\n", i)

		x := rand.Intn(5)

		switch x {
		case 0:
			fmt.Println("x is zero.")
		case 1:
			fmt.Println("x is one.")
		case 2:
			fmt.Println("x is two.")
		case 3:
			fmt.Println("x is three.")
		case 4:
			fmt.Println("x is four.")
		}

		fmt.Println()
	}
}
