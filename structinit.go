package main

import "fmt"

func main() {
   var(
      currency = struct{
         name, country string
         code int
      }{
         "USD", "United States",
         840,
      }
   ) 
   fmt.Println(currency.name)
   var(
      car = struct{make, model string}{make:"Ford", model:"F150"}
      node = struct{
         edges []string
         weight int
      }{
         edges: []string{"north", "south", "west"},
      }
   )
   fmt.Println(car.make)
   fmt.Println(node.edges)
}
