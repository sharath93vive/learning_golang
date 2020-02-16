package main

import "fmt"

func main() {
	tripMembers := map[string]float64{
		"sharath": 100,
		"sam":     50,
		"krangle": 60,
	}

	fmt.Println("People who paid money")
	for key, value := range tripMembers {
		fmt.Printf("%s : %0.2f\n", key, value)
	}

	//deleting a value using delete function

	delete(tripMembers, "krangle")
	fmt.Println(tripMembers)
}
