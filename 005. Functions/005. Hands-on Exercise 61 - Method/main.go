package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Printf("My name is %v and I am %v.\n", p.first, p.age)
}

func main() {
	p1 := person{
		first: "Pieter",
		age:   43,
	}

	p1.speak()
}
