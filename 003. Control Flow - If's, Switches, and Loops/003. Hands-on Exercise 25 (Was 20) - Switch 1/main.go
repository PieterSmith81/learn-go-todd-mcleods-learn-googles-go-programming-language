package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250) // Generate a random integer value between 0 and 249.

	fmt.Printf("The random value of x is %d.\n", x)

	switch {
	case x <= 100:
		fmt.Println("It is between 0 and 100.")
	case x >= 101 && x <= 200:
		fmt.Println("It is between 101 and 200.")
	default:
		fmt.Println("It is between 201 and 250.")
	}
}
