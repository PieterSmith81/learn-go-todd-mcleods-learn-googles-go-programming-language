package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	// Ranging over a map, displaying the key and value for each property
	for k, v := range m {
		fmt.Println(k, v)
	}
}
