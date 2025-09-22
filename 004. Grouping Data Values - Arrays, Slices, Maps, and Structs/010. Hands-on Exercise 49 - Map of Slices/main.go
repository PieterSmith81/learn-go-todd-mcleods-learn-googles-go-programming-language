package main

import "fmt"

func main() {
	// The map's key is of type string and its value is a slice of stings
	m := map[string][]string{
		`bond_james`:       {`shaken, not stirred`, `martinis`, `fast cars`},
		`moneypenny_jenny`: {`intelligence`, `literature`, `computer science`},
		`no_dr`:            {`cats`, `ice cream`, `sunsets`},
	}

	/* You can also define a map (or an slice, or array for that matter)
	using Go's builtin make() function, like so:

	m := make(map[string][]string)
	m[`bond_james`] = []string{`shaken, not stirred`, `martinis`, `fast cars`}
	m[`moneypenny_jenny`] = []string{`intelligence`, `literature`, `computer science`}
	m[`no_dr`] = []string{`cats`, `ice cream`, `sunsets`} */

	fmt.Printf("%#v\n\n", m)

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
