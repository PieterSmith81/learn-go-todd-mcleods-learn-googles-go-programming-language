package main

import "fmt"

type person struct {
	first string
}

func main() {
	// Value semantics
	p1 := person{
		first: "Pieter",
	}
	fmt.Println(p1)

	p1 = changeByValue(p1, "Ida")
	fmt.Println(p1)
	fmt.Println()

	// Pointer semantics
	p2 := person{
		first: "Mabel",
	}
	fmt.Println(p2)

	changeByPointer(&p2, "Willow")
	fmt.Println(p2)
}

func changeByValue(p person, newFirst string) person {
	p.first = newFirst
	return p
}

func changeByPointer(p *person, newFirst string) {
	p.first = newFirst
}
