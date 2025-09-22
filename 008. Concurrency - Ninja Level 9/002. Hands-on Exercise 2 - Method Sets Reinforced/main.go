package main

import "fmt"

type person struct{}

type human interface {
	speak()
}

func main() {
	p1 := person{}
	// p1.speak() // Works

	// You CAN pass a value of type *person into saySomething.
	saySomething(&p1) // Works

	// You CANNOT pass a value of type person into saySomething.
	// saySomething(p1) // Does not work!

	// Summary of the interface with values/pointers rules
	// Receivers		Values
	// ---------------------------------------------------
	// (t T)			T and *T
	// (t *T)			*T
}

func saySomething(h human) {
	h.speak()
}

func (p *person) speak() {
	fmt.Println("You shall not pass!")
}
