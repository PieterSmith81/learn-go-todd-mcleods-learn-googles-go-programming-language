package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	// Lookup values in a map based on the key
	age := m["James"]
	fmt.Println("The age of Bond:", age)

	age = m["Moneypenny"]
	fmt.Println("The age of Moneypenney:", age)

	age = m["Q"]
	fmt.Println("The age of Q:", age)

	// "comma ok" idiom examples
	if age, ok := m["James"]; ok {
		fmt.Println("The age of Bond (again):", age)
	} else {
		fmt.Println("The map does not contain the key you specified (i.e., there is no \"James\" lookup entry in the map).")
	}

	if age, ok := m["Q"]; ok {
		fmt.Println("The age of Q (again):", age)
	} else {
		fmt.Println("The map does not contain the key you specified (i.e., there is no \"Q\" lookup entry in the map).")
	}

}
