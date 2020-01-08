package main

import "fmt"

func main() {
   var val [100]int = [100]int{1,2,3,4,5}
   for _, i := range val {
      fmt.Println(i)
   }
   var truth = [256]bool{true}
   for _, j := range truth {
      fmt.Println(j)
   }
   var histogram = [5]map[string]int {
      map[string]int{"A":12, "B":5, "D":11},
      map[string]int{"man":1344, "women":844, "children":577},
   }
   for _, k := range histogram {
      fmt.Println(k)
   }
}
