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
tables := []map[string][]int {
         { 
               "age":{53, 13, 5, 55, 45, 62, 34, 7}, 
               "pay":{124, 66, 777, 531, 933, 231}, 
         }, 
    } 
   for _, i := range tables {
      fmt.Println(i)
   }
    graph  := [][][][]int{ 
         {{{44}, {3, 5}}, {{55, 12, 3}, {22, 4}}}, 
         {{{22, 12, 9, 19}, {7, 9}}, {{43, 0, 44, 12}, {7}}},     
    }
   for _, i := range graph {
      fmt.Println(i)
   }
}
