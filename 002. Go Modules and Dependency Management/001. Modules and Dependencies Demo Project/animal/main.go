/* Notes:
In this example, animal has a direct dependency on puppy, and puppy has a direct dependency on dog, as per the diagram below:
animal <-- puppy <-- dog

So, animal has a direct dependency on puppy, and an indirect dependency on dog.

Also, animal is a local Go module (not linked/hosted on GitHub at all).
puppy and dog are local Go modules that have been pushed to GitHub (so both are hosted on GitHub).
And we are using the GitHub-hosted puppy Go module as a dependency on this animal Go module - see the go.mod file in this folder (i.e. the /animal folder) for the dependency require statement. */

package main

import (
	"fmt"

	"github.com/PieterSmith81/puppy"
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)

	// Or you can also call the function in the imported dependency module "puppy" directly using nested function calls, like so:
	// fmt.Println(puppy.Bark())
	// fmt.Println(puppy.Barks())

	s3 := puppy.BigBark()
	s4 := puppy.BigBarks()
	fmt.Println(s3)
	fmt.Println(s4)
}
