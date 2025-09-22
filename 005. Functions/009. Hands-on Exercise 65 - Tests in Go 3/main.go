package main

import "fmt"

func main() {
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", subtract)
	fmt.Printf("%T\n", doMath)

	x := doMath(42, 16, add)
	fmt.Println(x)

	y := doMath(42, 16, subtract)
	fmt.Println(y)

	favMovie := favMovies(1)
	fmt.Println(favMovie)
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func favMovies(position int) string {
	switch position {
	case 1:
		return "The Matrix"
	case 2:
		return "Zoolander"
	case 3:
		return "Braveheart"
	default:
		return ""
	}
}
