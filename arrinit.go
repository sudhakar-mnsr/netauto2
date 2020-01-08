package main

import "fmt"

func main() {
   var val [100]int = [100]int{1,2,3,4,5}
   for _, i := range val {
      fmt.Println(i)
   }
}
