package main

import "fmt"

func main() {

var ids []string = []string{"fe225","ac144","3b12c"}
   for _, i := range ids {
      fmt.Println(i)
   }
vector := []float64{12.4, 44, 126, 2, 11.5}
   for _, i := range vector {
      fmt.Println(i)
   }
months := []string {
         "Jan", "Feb", "Mar", "Apr", 
         "May", "Jun", "Jul", "Aug", 
         "Sep", "Oct", "Nov", "Dec", 
         }
   for _, i := range months {
      fmt.Println(i)
   }
}
