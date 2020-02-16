package main

import "fmt"

func main() {

	ranks := map[string]int {"Gold":1, "silver":2, "bronze":3}
	for index, element := range ranks {
		fmt.Println(index, element)
	}

	// multiline literal mapping

	elements := map[string]string {
		"H":"Hydrogen",
		"Li":"Lithium",
	}
	fmt.Println("Printing without range function")
	fmt.Println(elements["Li"])
	fmt.Println(elements["H"])
	fmt.Println("Printing with range function")
	for _, element := range elements {
		fmt.Println(element)
	}

	// trying to access a key which was not assigned any value, will return zero or zero according to the perticular data type
	fmt.Println("empty value : ", elements["Ni"])
	
	// There is a confusion, whether zero is assigned to the key and its a default value

	// make use of boolean value called ok

	value, ok := elements["Ni"]
	fmt.Println(value, ok)
	elements["Ni"] = ""
	value, ok = elements["Ni"]
	fmt.Println(value, ok)
}
