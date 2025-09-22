package main

import "fmt"

func main() {
	i := 0

	// Infinite loop example (with a break) - increment.
	for {
		if i >= 20 {
			break
		}

		fmt.Println(i)
		i++
	}

	fmt.Println()

	j := 20
	// Infinite loop example (with a break) - decrement.
	for {
		if j < 0 {
			break
		}

		fmt.Println(j)
		j--
	}
}
