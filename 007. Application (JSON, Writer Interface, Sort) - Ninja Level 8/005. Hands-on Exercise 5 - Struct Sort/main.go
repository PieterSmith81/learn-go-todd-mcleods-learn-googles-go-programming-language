package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

// See the first example in the Overview section of the sort package documentation below
// for the example code / interface implementation that I based my solution here on:
// https://pkg.go.dev/sort#pkg-overview

// ByAge implements sort.Interface for []Person based on the Age field.
type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// ByLast implements sort.Interface for []Person based on the Last field.
type ByLast []user

func (l ByLast) Len() int           { return len(l) }
func (l ByLast) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ByLast) Less(i, j int) bool { return l[i].Last < l[j].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)
	fmt.Println()

	// your code goes here

	// Sorted by age
	sort.Sort(ByAge(users))
	fmt.Println(users)
	fmt.Println()

	// Sorted by last name
	sort.Sort(ByLast(users))
	fmt.Println(users)
	fmt.Println()

	// Sorted by last name (from previous sort above) and with sayings sorted too
	for _, user := range users {
		sort.Strings(user.Sayings)
	}
	fmt.Println(users)
	fmt.Println()

	// Printed in a "pleasant" way
	for i, user := range users {
		fmt.Printf("User %d:\n", i)
		fmt.Println("\tFirst name:", user.First)
		fmt.Println("\tLast name:", user.Last)
		fmt.Println("\tAge:", user.Age)
		fmt.Println("\tSayings:")

		for _, saying := range user.Sayings {
			fmt.Println("\t\t", saying)
		}

		fmt.Println()
	}
}
