package main

import "fmt"

func main() {
	i := 0

	// Similar to a while loop in other languages (with incrementation of the variable checked in the condition)
	for i < 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println()

	j := 20

	// Again, similar to a while loop in other languages (with decrementation of the variable checked in the condition)
	for j >= 0 {
		fmt.Println(j)
		j--
	}
}
