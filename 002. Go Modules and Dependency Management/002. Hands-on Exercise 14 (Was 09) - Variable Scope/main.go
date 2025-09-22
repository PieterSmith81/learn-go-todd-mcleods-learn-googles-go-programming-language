package main

import "fmt"

// Package level
var packageLevelVar = "I am a package level variable."

/* or
var packageLevelVar string = "I am a package level variable." */

const packageLevelConst = "I am a package level constant."

/* or
const packageLevelConst string= "I am a package level constant." */

func main() {
	// Block level
	blockLevelVar := "I am a block level variable."

	fmt.Println(packageLevelVar)
	fmt.Println(packageLevelConst)
	fmt.Println(blockLevelVar)
}
