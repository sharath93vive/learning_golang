package main

import (
	"fmt"
)

type part struct {
	description string
	count       int
}

type car struct {
	name     string
	topspeed float64
}

func main() {
	var posche car
	posche.name = "Poche 911 R"
	posche.topspeed = 323
	fmt.Printf("Car Name : %s\n", posche.name)
	fmt.Printf("TopSpeed : %0.2f\n", posche.topspeed)

	var bolts part
	bolts.description = "Bolt for car"
	bolts.count = 3
	fmt.Printf("part description : %s\n", bolts.description)
	fmt.Printf("Quantity : %d\n", bolts.count)
}
