package main

import "fmt"

var countToTen = func() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func main() {
	printSomething := func(message string) {
		fmt.Println(message)
	}

	printSomething("Let's Go!")

	countToTen()
}
