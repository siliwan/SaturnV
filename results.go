package main

import (
	"fmt"
)

// Final spped, and height
type Results struct {
	Steps int
	Stages []Flight
	Velocity float64
	Altitude float64
}

// All values are stored at the end of a stage.
type Flight struct {
	Velocity float64
	Altitude float64
	Weight float64
}

func results() {
	fmt.Println("- - - - -")
	fmt.Println("Results:")
	fmt.Println("- - - - -")

	fmt.Println("\n")

	fmt.Println("Rocket:", rocket.Name)
	fmt.Println("a")

}
