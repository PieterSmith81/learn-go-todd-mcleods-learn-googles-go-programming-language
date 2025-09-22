package main

import "fmt"

type customErr struct {
}

func main() {
	ce1 := customErr{}

	// The fact that the code line below doesn't throw an error
	// proves that ce is both a customErr struct as well as an "error interface".
	foo(ce1)
}

func (ce customErr) Error() string {
	return "I am a custom error implementing the error interface method signature!"
}

func foo(e error) {
	fmt.Println(e)
}
