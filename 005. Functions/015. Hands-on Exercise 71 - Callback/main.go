package main

import "fmt"

func main() {
	fmt.Println(printSquare(square, 2))
}

func square(n int) int {
	return n * n
}

func printSquare(f func(n int) int, x int) string {
	return fmt.Sprintf("The number %d squared is %d.", x, f(x))
}
