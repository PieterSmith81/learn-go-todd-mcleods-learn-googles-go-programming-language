package main

import "fmt"

func main() {
	// Multi-dimensional slice with composite literal value assignment
	x := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "I'm 008."}}

	fmt.Println(x)
	fmt.Println()

	// Outer loop
	for i, v := range x {
		fmt.Print(i, ". ")

		// Inner loop
		for i2, v2 := range v {
			fmt.Printf("\"%v (%v)\" ", v2, i2)
		}

		fmt.Println()
	}
}
