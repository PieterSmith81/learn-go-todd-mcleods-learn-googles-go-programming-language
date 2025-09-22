package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {

	wg.Add(2)

	go saySomething1()
	go saySomething2()

	wg.Wait()
}

func saySomething1() {
	fmt.Println("Say something 1 ran and completed.")
	wg.Done()
}

func saySomething2() {
	fmt.Println("Say something 2 ran and completed.")
	wg.Done()
}
