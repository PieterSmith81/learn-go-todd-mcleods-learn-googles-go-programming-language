package main

import "fmt"

func main() {
	x := make([]string, 0, 50)

	fmt.Println("Slice x's length on creation:", len(x))
	fmt.Println("Slice x's capacity on creation:", cap(x))
	fmt.Println()

	x = append(x, `Alabama`, `Alaska`, `Arizona`, `Arkansas`, `California`, `Colorado`, `Connecticut`, `Delaware`, `Florida`, `Georgia`, `Hawaii`, `Idaho`, `Illinois`, `Indiana`, `Iowa`, `Kansas`, `Kentucky`, `Louisiana`, `Maine`, `Maryland`, `Massachusetts`, `Michigan`, `Minnesota`, `Mississippi`, `Missouri`, `Montana`, `Nebraska`, `Nevada`, `New Hampshire`, `New Jersey`, `New Mexico`, `New York`, `North Carolina`, `North Dakota`, `Ohio`, `Oklahoma`, `Oregon`, `Pennsylvania`, `Rhode Island`, `South Carolina`, `South Dakota`, `Tennessee`, `Texas`, `Utah`, `Vermont`, `Virginia`, `Washington`, `West Virginia`, `Wisconsin`, `Wyoming`)

	fmt.Println("Slice x's length after appending values to it:", len(x))
	fmt.Println("Slice x's capacity after appending values to it:", cap(x))
	fmt.Println()

	fmt.Println("Index\tValue")
	for i := 0; i < len(x); i++ {
		fmt.Printf("%v\t%v\n", i, x[i])
	}
}
