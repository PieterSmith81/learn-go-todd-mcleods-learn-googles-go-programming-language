package main

import "fmt"

func main() {
	type engine struct {
		electric bool
	}

	type vehicle struct {
		engine
		make  string
		model string
		doors int
		color string
	}

	v1 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "Telsa",
		model: "Model S",
		doors: 5,
		color: "silver",
	}

	v2 := vehicle{
		engine: engine{
			electric: false,
		},
		make:  "Lamborghini",
		model: "Revuelto",
		doors: 2,
		color: "Orange",
	}

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println()

	fmt.Println(v1.make, v1.model, "- electric:", v1.electric)
	fmt.Println(v2.make, v2.model, "- electric:", v2.electric)
}
