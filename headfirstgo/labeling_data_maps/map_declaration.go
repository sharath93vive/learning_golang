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

}
