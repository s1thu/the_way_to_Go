package main

import "fmt"

type Shaper interface {
  Area() float32
}

type Square struct {
  side float32
}

func (sq *Square) Area() float32 {
  return sq.side * sq.side
}

func main1() {
  sq1 := new(Square)
  sq1.side = 5
  var areaIntf Shaper
  areaIntf = sq1
  // shorter, without separate declaration:
  // areaIntf := Shaper(sq1)
  fmt.Printf("The square has area: %f\n", areaIntf.Area())
}