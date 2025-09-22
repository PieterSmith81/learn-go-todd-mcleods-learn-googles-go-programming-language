package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}

func main() {
	fmt.Printf("The value of a is %v and its type is %T.\n", a, a)
	fmt.Printf("The value of b is %v and its type is %T.\n", b, b)
	fmt.Printf("The value of c is %v and its type is %T.\n", c, c)
	fmt.Printf("The value of d is %v and its type is %T.\n", d, d)
	fmt.Println()

	fmt.Printf("The (dereferenced) value of a is: %v.\n", *a)
	fmt.Printf("The (dereferenced) value of b is: %v.\n", *b)
	fmt.Printf("The (dereferenced) value of c is: %v.\n", *c)
	fmt.Printf("The (dereferenced) value of d is: %v.\n", *d)
}
