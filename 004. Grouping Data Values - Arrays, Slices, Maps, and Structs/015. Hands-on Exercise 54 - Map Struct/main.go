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

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, v := range m {
		fmt.Println()
		for i, v2 := range v.favouriteIceCream {
			fmt.Printf("%v. %v's favourite is %v.\n", i, v.firstName, v2)
		}
	}
}
