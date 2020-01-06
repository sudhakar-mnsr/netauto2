package main

import "fmt"

func avg(nums ...float64) float64 {
   n := len(nums)
   t := 0.0
   for _, v := range nums {
      t += v
   }
   return t / float64(n)
}

func main() {
   fmt.Printf("avg([1, 2.5, 3.75]) =%.2f\n", avg(1, 2.5, 3.75))
}
