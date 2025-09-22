package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	y := foo(xi...)
	fmt.Println(y)

	z := bar(xi)
	fmt.Println(z)
}

func foo(i ...int) int {
	total := 0

	for _, v := range i {
		total += v
	}

	return total
}

func bar(xi []int) int {
	total := 0

	for _, v := range xi {
		total += v
	}

	return total
}
