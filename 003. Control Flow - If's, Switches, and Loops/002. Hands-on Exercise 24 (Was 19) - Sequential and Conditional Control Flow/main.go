package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250) // Generate a random integer value between 0 and 249.

	fmt.Printf("The random value of x is %d.\n", x)

	if x <= 100 {
		fmt.Println("It is between 0 and 100.")
	} else if x >= 101 && x <= 200 {
		fmt.Println("It is between 101 and 200.")
	} else {
		fmt.Println("It is between 201 and 250.")
	}

	/* Regarding, inclusive or exclusive – does rand.Intn()
	- include zero in its output?
	- include 250 in its output?
	- show this in code using the numbers 0 - 3 */
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))

	/* Answer, it is exclusive, which can also be seen in the "half-open" reference in the official Go documentation here:
	https://pkg.go.dev/math/rand#Intn */
}
