package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// The code below creates a race condition.
	// You can prove this by running the code using the "-race" flag
	// with the "go run" command, i.e., by running:
	// go run -race main.go
	count := 0
	const numOfGoRoutines = 100

	wg := sync.WaitGroup{}
	wg.Add(numOfGoRoutines)

	for i := 0; i < numOfGoRoutines; i++ {
		go func() {
			v := count
			// Gosched yields the processor, allowing other goroutines to run. It does not
			// suspend the current goroutine, so execution resumes automatically.
			runtime.Gosched()
			v++
			count = v
			fmt.Println("Count:", count)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Final Count:", count)
}
