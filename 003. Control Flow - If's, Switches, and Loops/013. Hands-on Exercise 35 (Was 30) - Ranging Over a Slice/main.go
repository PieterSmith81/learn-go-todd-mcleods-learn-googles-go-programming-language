package main

import "fmt"

func main() {
	xi := []int{42, 43, 44, 45, 46, 47}

	// Ranging over a slice (printing the index and value)
	for i, v := range xi {
		fmt.Println(i, v)
	}
}
