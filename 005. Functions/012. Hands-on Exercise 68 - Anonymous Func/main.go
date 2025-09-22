package main

import "fmt"

func main() {
	func(message string) {
		fmt.Println(message)
	}("Let's Go!")
}
