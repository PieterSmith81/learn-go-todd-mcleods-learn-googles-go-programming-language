package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)

	// Basic error handling
	if err != nil {
		log.Fatalln("An error occurred converting your data to JSON. Error details:", err)
	}

	fmt.Println(string(bs))
}
