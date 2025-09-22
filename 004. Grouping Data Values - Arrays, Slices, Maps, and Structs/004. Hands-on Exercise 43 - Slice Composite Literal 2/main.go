package main

import "fmt"

func main() {
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println("Index\tValue\tType")

	for index, value := range slice {
		fmt.Printf("%v\t%v\t%T\n", index, value, value)
	}
}
