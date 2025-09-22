package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		// Only print i when it's an odd number
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}
