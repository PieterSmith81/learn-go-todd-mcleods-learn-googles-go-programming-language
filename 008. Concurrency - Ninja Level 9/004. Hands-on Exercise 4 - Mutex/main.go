package main

import (
	"fmt"
	"sync"
)

func main() {
	// Fixed the race condition from the previous exercise using a Mutex.
	// You can prove this by running the code using the "-race" flag
	// with the "go run" command, i.e., by running:
	// go run -race main.go
	count := 0
	const numOfGoRoutines = 100

	wg := sync.WaitGroup{}
	wg.Add(numOfGoRoutines)

	var m sync.Mutex

	for i := 0; i < numOfGoRoutines; i++ {
		go func() {
			m.Lock()
			v := count
			v++
			count = v
			fmt.Println("Count:", count)
			m.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Final Count:", count)
}
