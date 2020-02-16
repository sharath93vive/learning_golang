package main

import "fmt"

type subscriber struct {
  name string
  rate float64
  active bool
}

func main() {

  var subscriber1 subscriber
  subscriber1.name = "sharath"
  subscriber1.rate = 1000
  subscriber1.active = true

  fmt.Println("%#v", subscriber1)
}
