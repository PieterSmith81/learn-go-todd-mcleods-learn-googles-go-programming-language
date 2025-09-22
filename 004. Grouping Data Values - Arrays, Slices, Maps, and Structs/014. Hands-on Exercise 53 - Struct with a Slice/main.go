package main

import "fmt"

func main() {
	type person struct {
		firstName         string
		lastName          string
		favouriteIceCream []string
	}

	p1 := person{
		firstName:         "James",
		lastName:          "Bond",
		favouriteIceCream: []string{"Chocolate", "Pistachio", "Vanilla"},
	}

	p2 := person{
		firstName:         "Eve",
		lastName:          "Moneypenny",
		favouriteIceCream: []string{"Strawberry", "Mango", "Hazelnut"},
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println()

	for i, v := range p1.favouriteIceCream {
		fmt.Printf("%v. %v's favourite is %v.\n", i, p1.firstName, v)
	}

	fmt.Println()

	for i, v := range p2.favouriteIceCream {
		fmt.Printf("%v. %v's favourite is %v.\n", i, p2.firstName, v)
	}
}
