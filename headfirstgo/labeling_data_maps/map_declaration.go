package main

import "fmt"

func main() {
	// Declaring map
	var myMap map[string]float64
	// create a map using make function
	myMap = make(map[string]float64)
	myMap["height"] = 6.5
	myMap["weight"] = 70.23

	fmt.Println(myMap)

	// short variable declaration
	ranks := make(map[string]int)
	ranks["glod"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3

	fmt.Println(ranks)

	// map with string as key and string as value

	elements := make(map[string]string)
	elements["Li"] = "Lithium"
	elements["H"] = "Hydrogen"

	fmt.Println(elements["Li"])
	fmt.Println(elements["H"])

	//map with integer as a key and boolean as value

	isPrime := make(map[int]bool)
	isPrime[4] = false
	isPrime[7] = true

	fmt.Println(isPrime)

	// map literals

	someContent := map[string]float64{"a": 1, "b": 2, "c": 3}
	fmt.Println(someContent)

}
