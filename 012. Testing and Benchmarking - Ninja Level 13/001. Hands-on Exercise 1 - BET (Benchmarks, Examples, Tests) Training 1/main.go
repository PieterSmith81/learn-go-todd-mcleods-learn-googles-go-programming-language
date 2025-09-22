package main

import (
	"fmt"
	"myApp/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	willow := canine{
		name: "Willow",
		age:  dog.Years(5),
	}
	fmt.Println(willow)
	fmt.Println(dog.YearsTwo(20))
}
