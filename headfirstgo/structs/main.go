package main

import "fmt"

var myStruct struct {
	number float64
	word   string
	toggle bool
}

func main() {
	fmt.Printf("%#v\n", myStruct)
	myStruct.number = 10.45
	myStruct.word = "Hello"
	myStruct.toggle = true

	fmt.Println(myStruct.number, myStruct.word, myStruct.toggle)
}
