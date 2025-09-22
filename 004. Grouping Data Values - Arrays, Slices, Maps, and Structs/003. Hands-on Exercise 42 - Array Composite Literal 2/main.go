package main

import "fmt"

func main() {
	array := [5]int{0, 1, 2, 3, 4}

	fmt.Println("Index\tValue\tType")

	for index, value := range array {
		fmt.Printf("%v\t%v\t%T\n", index, value, value)
	}
}
