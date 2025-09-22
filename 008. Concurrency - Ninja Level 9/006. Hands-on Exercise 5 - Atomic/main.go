package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// Fixed the race condition from the previous exercise using Atomic.
	// You can prove this by running the code using the "-race" flag
	// with the "go run" command, i.e., by running:
	// go run -race main.go
	var count int64
	const numOfGoRoutines = 100

	wg := sync.WaitGroup{}
	wg.Add(numOfGoRoutines)

	for i := 0; i < numOfGoRoutines; i++ {
		go func() {
			atomic.AddInt64(&count, 1)
			fmt.Println("Count:", atomic.LoadInt64(&count))
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Final Count:", count)
}
