package main

import (
	"fmt"
	"math/rand"
	"time"
)
func randomNumberGen(min float32, max float32) float32 {	
	return min + (max-min)*rand.Float32()
}

type Stock struct {
  name string
  price float32                                             
}

func (stk *Stock) updateStock() {
  change := randomNumberGen(-10000, 10000)
  stk.price += change
}

func displayMarket(market []Stock) {
  for i := 0; i < len(market); i++ {
    fmt.Println(market[i])
  }
}

func main() {
  rand.Seed(time.Now().UnixNano())

  stockMarket := []Stock{{"GOOG", 2313.50}, {"AAPL", 157.28}, {"FB", 203.77}, {"TWTR", 50.00}}

  displayMarket(stockMarket)
  fmt.Println()

  stockMarket[0].updateStock()
  stockMarket[1].updateStock()
  stockMarket[2].updateStock()
  stockMarket[3].updateStock()

  displayMarket(stockMarket)
  fmt.Println()

}
