package main

import "fmt"

type part struct {
  description string
  count int
}

func (p part) getInfo() {
  fmt.Println(p.description, p.count)
}

func minimunOrder (description string, count int) (part){
  var tempPart part
  tempPart.description = "Default description"
  tempPart.count = 10

  return tempPart
}

func main() {

var part1 part
part1 = minimunOrder("default value", 100)
part1.getInfo()
}
