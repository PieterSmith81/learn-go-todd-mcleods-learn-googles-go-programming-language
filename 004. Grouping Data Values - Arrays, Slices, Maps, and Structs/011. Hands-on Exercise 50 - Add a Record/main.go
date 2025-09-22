package main

import "fmt"

func main() {
	// The map's key is of type string and its value is a slice of stings
	m := map[string][]string{
		`bond_james`:       {`shaken, not stirred`, `martinis`, `fast cars`},
		`moneypenny_jenny`: {`intelligence`, `literature`, `computer science`},
		`no_dr`:            {`cats`, `ice cream`, `sunsets`},
	}

	// Add a record to the map
	m[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}

	// Range over the map
	for k, v := range m {
		fmt.Print(k, ":\n")
		// Range over the slice (the map's values, which are slices)
		for i, v2 := range v {
			fmt.Printf("\t%v. %v\n", i, v2)
		}
		fmt.Println()
	}
}
