package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// write your error code here

		// e := errors.New("More coffee needed!") // Does the same as the code line below, but fmt.Errorf() is, arguably, a more elegant solution.
		e := fmt.Errorf("More coffee needed! Value was: %v.", f)

		return 0, sqrtError{"50.2289 N", "99.4656 W", e}
	}

	// fictitious square root calculation/return
	return 42, nil
}
