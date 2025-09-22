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

	bs, err := toJSON(p1)

	// My error handling implementation
	if err != nil {
		log.Fatalln(err, "\n\nFatal application error. Exiting application.")
	}

	fmt.Println(string(bs))
}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)

	// My error handling implementation
	if err != nil {
		return nil, fmt.Errorf("An error occurred converting your data to JSON.\n\nError details:\n%v", err)
	}

	return bs, nil
}
