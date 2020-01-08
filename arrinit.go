package main

import "fmt"

func main() {
   var val [100]int = [100]int{1,2,3,4,5}
   for _, i := range val {
      fmt.Println(i)
   }
   var truth = [256]bool{true}
   for _, i := range truth {
      fmt.Println(i)
   }
   var histogram = [5]map[string]int {
      map[string]int{"A":12, "B":5, "D":11},
      map[string]int{"man":1344, "women":844, "children":577},
   }
   for _, i := range histogram {
      fmt.Println(i)
   }
   var board = [4][2]int{ 
      {33, 23}, 
      {62, 2}, 
      {23, 4}, 
      {51, 88}, 
   }
   for _, i := range board {
      fmt.Println(i)
   }
}
