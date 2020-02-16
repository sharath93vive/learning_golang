package main

import "fmt"

func main() {

    check_passing("alma")
    check_passing("rohit")
    check_passing("ken")
}

func check_passing(name string) {
    marks := map[string]float64{"alma":0, "rohit":76}
    value, ok := marks[name]
    if !ok {
        fmt.Println("No marks was initilized")

    } else {

    if value > 60 {
	    fmt.Printf("You are passing %s with %0.2f makrs\n", name, value)
    } else {
	    fmt.Println("You are failing")
    }
    }
}
